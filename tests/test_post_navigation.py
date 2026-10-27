"""Browser checks: install Playwright Chromium, then run unittest discovery."""
import functools
import http.server
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import threading
import unittest

from playwright.sync_api import sync_playwright

ROOT = Path(__file__).resolve().parents[1]


class PostNavigationTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.temp = tempfile.TemporaryDirectory(prefix="blog-navigation-")
        root = Path(cls.temp.name)
        shutil.copytree(ROOT / "themes", root / "themes")
        (root / "hugo.toml").write_text('baseURL = "/"\ntheme = "kukudely"\n', encoding="utf-8")
        posts = root / "content/posts"
        posts.mkdir(parents=True)
        body = "\n".join(f"## Chapter {i}\n\n" + "Reading text. " * 100 for i in range(40))
        (posts / "long.md").write_text('+++\ntitle = "Long article"\n+++\n' + body, encoding="utf-8")
        (posts / "short.md").write_text('+++\ntitle = "Short article"\n+++\nNo headings.', encoding="utf-8")
        subprocess.run([os.environ.get("HUGO_BIN", "hugo"), "--source", str(root)], check=True, capture_output=True)
        cls.server = http.server.ThreadingHTTPServer(("127.0.0.1", 0), functools.partial(
            http.server.SimpleHTTPRequestHandler, directory=str(root / "public")))
        threading.Thread(target=cls.server.serve_forever, daemon=True).start()
        cls.url = f"http://127.0.0.1:{cls.server.server_port}"
        cls.playwright = sync_playwright().start()
        cls.browser = cls.playwright.chromium.launch(headless=True)

    @classmethod
    def tearDownClass(cls):
        cls.browser.close()
        cls.playwright.stop()
        cls.server.shutdown()
        cls.server.server_close()
        cls.temp.cleanup()

    def setUp(self):
        self.page = self.browser.new_page(viewport={"width": 1440, "height": 900})
        self.errors = []
        self.page.on("pageerror", lambda error: self.errors.append(str(error)))
        self.page.goto(self.url + "/posts/long/", wait_until="networkidle")

    def tearDown(self):
        self.page.close()
        self.assertEqual(self.errors, [])

    def test_wide_layout_keeps_posts_reachable_with_long_toc(self):
        post_link = self.page.get_by_role("link", name="Short article")
        article = self.page.locator("article").bounding_box()
        link = post_link.bounding_box()
        self.assertLess(link["x"] + link["width"], article["x"])
        self.assertLess(link["y"], 900)
        self.page.evaluate("window.scrollTo(0, 3000)")
        self.page.wait_for_timeout(700)
        self.assertGreater(post_link.bounding_box()["y"], 0)
        post_link.click()
        self.assertTrue(self.page.url.endswith("/posts/short/"))

    def test_medium_tabs_and_keyboard(self):
        self.page.set_viewport_size({"width": 1100, "height": 800})
        posts = self.page.get_by_role("tab", name="Posts", exact=True)
        posts.click()
        self.page.get_by_role("link", name="Short article").wait_for(state="visible")
        posts.press("ArrowLeft")
        self.assertEqual(self.page.get_by_role("tab", name="本文目录").get_attribute("aria-selected"), "true")

    def test_mobile_navigation_stays_available_and_closes_after_jump(self):
        self.page.set_viewport_size({"width": 390, "height": 844})
        self.page.evaluate("window.scrollTo(0, 2000)")
        self.page.wait_for_timeout(700)
        toc = self.page.get_by_role("button", name="本文目录", exact=True)
        toc.click()
        self.page.get_by_role("link", name="Chapter 20", exact=True).click()
        self.assertEqual(toc.get_attribute("aria-expanded"), "false")
        self.page.wait_for_timeout(800)
        heading = self.page.locator("#chapter-20").bounding_box()
        self.assertGreaterEqual(heading["y"], toc.bounding_box()["y"] + toc.bounding_box()["height"])
        posts = self.page.get_by_role("button", name="Posts", exact=True)
        posts.click()
        posts.press("Escape")
        self.assertEqual(posts.get_attribute("aria-expanded"), "false")
        self.assertTrue(self.page.evaluate("document.documentElement.scrollWidth <= innerWidth"))
        posts.click()
        self.page.get_by_role("link", name="Short article").click()
        self.assertTrue(self.page.url.endswith("/posts/short/"))
        self.assertEqual(self.page.get_by_role("button", name="本文目录", exact=True).count(), 0)
