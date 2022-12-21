from factory import Factory, Link, Page, Tray


class DivLink(Link):
    def __init__(self, caption: str, url: str) -> None:
        super().__init__(caption, url)

    def makeHTML(self) -> str:
        return (
            '<div class="LINK"><a href="'
            + self.url
            + '">'
            + self.caption
            + "</a></div>\n"
        )


class DivTray(Tray):
    def __init__(self, caption) -> None:
        self.__string = ""
        super().__init__(caption)

    def makeHTML(self) -> str:
        self.__string += "<p><b>"
        self.__string += self.caption
        self.__string += "</b></p>\n"
        self.__string += '<div class="TRAY">'
        for item in self.tray:
            self.__string += item.makeHTML()
        self.__string += "</div>\n"
        return self.__string


class DivPage(Page):
    def __init__(self, title: str, author: str) -> None:
        self.__string = ""
        super().__init__(title, author)

    def makeHTML(self) -> str:
        self.__string += "<!DOCTYPE html>\n"
        self.__string += "<html><head><title>"
        self.__string += self.title
        self.__string += "</title><style>\n"
        self.__string += (
            "div.TRAY { padding:0.5em; margin-left:5em; border:1px solid black; }\n"
        )
        self.__string += "div.LINK { padding:0.5em; background-color: lightgray; }\n"
        self.__string += "</style></head><body>\n"
        self.__string += "<h1>\n"
        self.__string += self.title
        self.__string += "</h1>\n"
        for item in self.content:
            self.__string += item.makeHTML()
        self.__string += "<hr><address>"
        self.__string += self.author
        self.__string += "</address>\n"
        self.__string += "</body></html>\n"

        return self.__string


class DivFactory(Factory):
    def create_link(self, caption: str, url: str) -> Link:
        return DivLink(caption, url)

    def create_tray(self, caption: str) -> Tray:
        return DivTray(caption)

    def create_page(self, title: str, author: str) -> Page:
        return DivPage(title, author)
