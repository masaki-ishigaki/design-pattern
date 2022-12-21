import sys

from factory import Factory


def main() -> None:
    filename = sys.argv[1]
    classname = sys.argv[2]

    factory = Factory.get_factory(classname)

    # Blog
    blog1 = factory.create_link("Blog 1", "https://example.com/blog1")
    blog2 = factory.create_link("Blog 2", "https://example.com/blog2")
    blog3 = factory.create_link("Blog 3", "https://example.com/blog3")

    blog_tray = factory.create_tray("Blog Site")
    blog_tray.add(blog1)
    blog_tray.add(blog2)
    blog_tray.add(blog3)

    # News
    news1 = factory.create_link("News 1", "https://example.com/news1")
    news2 = factory.create_link("News 2", "https://example.com/news2")
    news3 = factory.create_tray("News 3")
    news3.add(factory.create_link("News 3 (US)", "https://exmaple.com/news3us"))
    news3.add(factory.create_link("News 3 (JP)", "https://exmaple.com/news3jp"))

    news_tray = factory.create_tray("News Site")
    news_tray.add(news1)
    news_tray.add(news2)
    news_tray.add(news3)

    # Page
    page = factory.create_page("Blog and News", "Hiroshi Yuki")
    page.add(blog_tray)
    page.add(news_tray)

    page.output(filename)


if __name__ == "__main__":
    main()
