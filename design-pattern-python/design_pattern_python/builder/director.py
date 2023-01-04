from builder import Builder


class Director:
    def __init__(self, builder: Builder):
        self.__builder = builder

    def constrcut(self):
        self.__builder.make_title("Greeting")
        self.__builder.make_string("一般的なあいさつ")
        self.__builder.make_items(
            [
                "How are you?",
                "Hello.",
                "Hi.",
            ]
        )
        self.__builder.make_string("時間帯に応じたあいさつ")
        self.__builder.make_items(
            [
                "Good morning.",
                "Good afternonn.",
                "Good evening.",
            ]
        )
        self.__builder.close()


if __name__ == "__main__":
    import sys

    builder_type = sys.argv[1]
    if builder_type == "text":
        from text_builder import TextBuilder

        text_builder = TextBuilder()
        director = Director(text_builder)
        director.constrcut()
        print(text_builder.get_text_result())
    elif builder_type == "html":
        from html_builder import HTMLBuilder

        html_builder = HTMLBuilder()
        director = Director(html_builder)
        director.constrcut()
        print(html_builder.get_html_result())
    else:
        exit()
