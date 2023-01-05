from typing import Dict, TextIO


class Database:
    def __new__(cls):
        raise NotImplementedError("Cannot generate instance by constructor")

    @classmethod
    def get_properties(cls, dbname: str) -> Dict[str, str]:
        filename = dbname + ".txt"
        properties: Dict[str, str] = {}
        f: TextIO

        with open(filename, "r") as f:
            for line in f.readlines():
                prop = line.split("=")
                properties[prop[0]] = prop[1].rstrip()

        return properties


class HtmlWriter:
    def __init__(self, writer: TextIO) -> None:
        self.__writer = writer

    # タイトルの出力
    def title(self, title: str) -> None:
        self.__writer.write("<!DOCTYPE html")
        self.__writer.write("<html>")
        self.__writer.write("<head>")
        self.__writer.write(f"<title>{title}</title>")
        self.__writer.write("</head>")
        self.__writer.write("<body>")
        self.__writer.write("\n")
        self.__writer.write(f"<h1>{title}</h1>")
        self.__writer.write("\n")

    # 段落の出力
    def paragraph(self, msg: str) -> None:
        self.__writer.write(f"<p>{msg}</p>")
        self.__writer.write("\n")

    # リンクの出力
    def link(self, href: str, caption: str) -> None:
        self.paragraph(f'<a href="{href}">{caption}</a>')

    # メールアドレスの出力
    def mailto(self, mailaddr: str, username: str) -> None:
        self.link(f"mailto: {mailaddr}", username)

    # 閉じる
    def close(self) -> None:
        self.__writer.write("</body>")
        self.__writer.write("</html>")
        self.__writer.write("\n")
        self.__writer.close()
