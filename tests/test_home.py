"""Build isolated Hugo fixtures: HUGO_BIN=/path/to/hugo python -m unittest discover -s tests."""
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest
from html.parser import HTMLParser


ROOT = Path(__file__).resolve().parents[1]


class HomeParser(HTMLParser):
    def __init__(self, html):
        super().__init__()
        self.stats = {}
        self.nodes = {}
        self.edges = set()
        self.coverage = {}
        self.feed(html)

    def handle_starttag(self, tag, attrs):
        attrs = dict(attrs)
        if "data-stat" in attrs:
            self.stats[attrs["data-stat"]] = int(attrs["data-count"])
        if "data-topic" in attrs:
            self.nodes[attrs["data-topic"]] = attrs["href"]
        if "data-from" in attrs:
            self.edges.add(frozenset((attrs["data-from"], attrs["data-to"])))
        if "data-map-posts" in attrs:
            self.coverage = {
                "posts": int(attrs["data-map-posts"]),
                "covered": int(attrs["data-map-covered"]),
            }


class HomeBuildTest(unittest.TestCase):
    def build_home(self, articles, config=""):
        with tempfile.TemporaryDirectory(prefix="hugo-home-") as temp:
            root = Path(temp)
            shutil.copytree(ROOT / "themes", root / "themes")
            (root / "hugo.toml").write_text(
                'baseURL = "https://example.org/"\ntheme = "kukudely"\n' + config, encoding="utf-8")
            for name, metadata in articles.items():
                path = root / "content" / f"{name}.md"
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text(f'+++\ntitle = "{name}"\n{metadata}\n+++\nExample body.\n', encoding="utf-8")
            result = subprocess.run(
                [os.environ.get("HUGO_BIN", "hugo"), "--source", str(root)],
                capture_output=True, text=True, encoding="utf-8")
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            return HomeParser((root / "public/index.html").read_text(encoding="utf-8"))

    def test_counts_and_connections_only_use_published_posts(self):
        home = self.build_home({
            "posts/one": 'date = "2020-01-01"\ntags = ["Alpha", "Beta"]',
            "posts/two": 'date = "2020-01-02"\ntags = ["Beta", "Gamma"]',
            "posts/draft": 'draft = true\ntags = ["Secret"]',
            "posts/future": 'date = "2099-01-01"\ntags = ["Future"]',
            "about/index": 'tags = ["Profile"]',
        })
        self.assertEqual(home.stats, {"posts": 2, "topics": 3})
        self.assertEqual(set(home.nodes), {"Alpha", "Beta", "Gamma"})
        self.assertEqual(home.nodes["Alpha"], "/tags/alpha/")
        self.assertEqual(home.edges, {frozenset(("Alpha", "Beta")), frozenset(("Beta", "Gamma"))})

    def test_empty_site_has_zero_stats_and_no_graph_links(self):
        home = self.build_home({"about/index": 'tags = ["Profile"]'})
        self.assertEqual(home.stats, {"posts": 0, "topics": 0})
        self.assertEqual(home.nodes, {})
        self.assertEqual(home.edges, set())

    def test_graph_is_bounded_but_overview_counts_all_topics(self):
        home = self.build_home({
            "posts/many": 'date = "2020-01-01"\ntags = ["a", "b", "c", "d", "e", "f", "g", "h"]',
        })
        self.assertEqual(home.stats, {"posts": 1, "topics": 8})
        self.assertEqual(len(home.nodes), 6)

    def test_recent_topics_win_when_article_counts_tie(self):
        home = self.build_home({
            "posts/older": 'date = "2020-01-01"\ntags = ["a", "b", "c", "d", "e", "f"]',
            "posts/newer": 'date = "2020-01-02"\ntags = ["Fresh"]',
        })
        self.assertEqual(len(home.nodes), 6)
        self.assertIn("Fresh", home.nodes)

    def test_low_frequency_topics_keep_all_articles_represented(self):
        home = self.build_home({
            "posts/common-one": 'date = "2020-01-01"\ntags = ["a", "b", "c", "d", "e", "f"]',
            "posts/common-two": 'date = "2020-01-02"\ntags = ["a", "b", "c", "d", "e", "f"]',
            "posts/database": 'date = "2019-01-01"\ntags = ["MySQL"]',
            "posts/site": 'date = "2021-01-01"\ntags = ["Hugo"]',
        })
        self.assertIn("MySQL", home.nodes)
        self.assertIn("Hugo", home.nodes)
        self.assertEqual(home.nodes["MySQL"], "/tags/mysql/")
        self.assertEqual(home.coverage, {"posts": 4, "covered": 4})
        self.assertEqual(len(home.nodes), 6)
        self.assertFalse(any("MySQL" in edge or "Hugo" in edge for edge in home.edges))

    def test_independent_topics_expand_to_cover_every_article(self):
        home = self.build_home({
            f"posts/{topic}": f'date = "2020-01-01"\ntags = ["{topic}"]'
            for topic in ("a", "b", "c", "d", "e", "f", "g", "h")
        })
        self.assertEqual(set(home.nodes), {"A", "B", "C", "D", "E", "F", "G", "H"})
        self.assertEqual(home.coverage, {"posts": 8, "covered": 8})
        self.assertEqual(home.edges, set())

    def test_untagged_post_remains_reachable(self):
        home = self.build_home({
            "posts/untagged": 'date = "2020-01-01"',
            "posts/tagged": 'date = "2020-01-02"\ntags = ["Alpha"]',
        })
        self.assertIn("posts/untagged", home.nodes)
        self.assertEqual(home.nodes["posts/untagged"], "/posts/untagged/")
        self.assertEqual(home.coverage, {"posts": 2, "covered": 2})

    def test_excluded_topics_match_normalized_taxonomy_titles(self):
        home = self.build_home({
            "posts/welcome": 'date = "2020-01-01"\ntags = ["hello", "hugo"]',
        }, '[params.home]\ngraphExcludeTags = ["hello"]\n')
        self.assertEqual(set(home.nodes), {"Hugo"})
        self.assertEqual(home.stats, {"posts": 1, "topics": 2})
        self.assertEqual(home.coverage, {"posts": 1, "covered": 1})


if __name__ == "__main__":
    unittest.main()
