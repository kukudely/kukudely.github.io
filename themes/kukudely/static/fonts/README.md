# 首页字体

- Newsreader：英文主标题，500 字重，支持自动光学尺寸。
- Noto Sans SC：中文与界面文字，400–600 可变字重。
- 来源：Google Fonts CSS API，下载日期 2026-10-26。
- 请求：`https://fonts.googleapis.com/css2?family=Newsreader:opsz,wght@6..72,500&family=Noto+Sans+SC:wght@400..600&display=swap`
- 两套字体均保留各自的 SIL Open Font License，见同目录的 `*-OFL.txt`。

`fonts.css` 保留官方 `unicode-range` 分片定义。浏览器按页面实际使用的字符下载所需 WOFF2 文件；字体与网站同源加载，不依赖访问者连接外部字体服务。更新时需保留完整分片覆盖，以支持新文章中的文字。
