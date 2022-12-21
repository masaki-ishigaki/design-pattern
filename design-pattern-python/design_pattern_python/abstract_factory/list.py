from factory import Factory, Link, Page, Tray


class ListLink(Link):
    def __init__(self, caption: str, url: str) -> None:
        super().__init__(caption, url)

    def makeHTML(self) -> str:
        return '  <li><a href="' + self.url + '">' + self.caption + "</a></li>\n"


class ListTray(Tray):
    def __init__(self, caption) -> None:
        self.__string = ""
        super().__init__(caption)

    def makeHTML(self) -> str:
        self.__string += "<li>\n"
        self.__string += self.caption
        self.__string += "\n<ul>\n"
        for item in self.tray:
            self.__string += item.makeHTML()
        self.__string += "</ul>\n"
        self.__string += "</li>\n"
        return self.__string


class ListPage(Page):
    def __init__(self, title: str, author: str) -> None:
        self.__string = ""
        super().__init__(title, author)

    def makeHTML(self) -> str:
        self.__string += "<!DOCTYPE html>\n"
        self.__string += "<html><head><title>"
        self.__string += self.title
        self.__string += "</title></head>\n"
        self.__string += "<body>\n"
        self.__string += "<h1>\n"
        self.__string += self.title
        self.__string += "</h1>\n"
        self.__string += "<ul>\n"
        for item in self.content:
            self.__string += item.makeHTML()
        self.__string += "</ul>\n"
        self.__string += "<hr><address>"
        self.__string += self.author
        self.__string += "</address>\n"
        self.__string += "</body></html>\n"

        return self.__string


class ListFactory(Factory):
    def create_link(self, caption: str, url: str) -> Link:
        return ListLink(caption, url)

    def create_tray(self, caption: str) -> Tray:
        return ListTray(caption)

    def create_page(self, title: str, author: str) -> Page:
        return ListPage(title, author)
