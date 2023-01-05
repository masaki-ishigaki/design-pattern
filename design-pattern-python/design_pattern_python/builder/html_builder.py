from typing import List

from builder import Builder, StringBuilder


class HTMLBuilder(Builder):
    def __init__(self) -> None:
        self.__filename = "untitled.html"
        self.__sb = StringBuilder()

    def make_title(self, title: str) -> None:
        self.__filename = title + ".html"
        self.__sb.append("<!DOCTYPE html> \n")
        self.__sb.append("<html>\n")
        self.__sb.append("<head><title>")
        self.__sb.append(title)
        self.__sb.append("</title></head>\n")
        self.__sb.append("<body>\n")
        self.__sb.append("<h1>")
        self.__sb.append(title)
        self.__sb.append("</h1>\n\n")

    def make_string(self, string: str) -> None:
        self.__sb.append("<p>")
        self.__sb.append(string)
        self.__sb.append("</p>\n\n")

    def make_items(self, items: List[str]) -> None:
        self.__sb.append("<ul>")
        for item in items:
            self.__sb.append("<li>")
            self.__sb.append(item)
            self.__sb.append("</li>\n")

        self.__sb.append("\n")

    def close(self) -> None:
        self.__sb.append("</body>")
        self.__sb.append("</html>\n")

        with open(self.__filename, "w") as f:
            f.write(self.__sb.string)

    def get_html_result(self) -> str:
        return self.__filename
