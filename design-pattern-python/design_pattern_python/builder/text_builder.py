from typing import List

from builder import Builder, StringBuilder


class TextBuilder(Builder):
    def __init__(self) -> None:
        self.__sb = StringBuilder()

    def make_title(self, title: str) -> None:
        self.__sb.append("======================================\n")
        self.__sb.append("『")
        self.__sb.append(title)
        self.__sb.append("』\n\n")

    def make_string(self, string: str) -> None:
        self.__sb.append("■")
        self.__sb.append(string)
        self.__sb.append("\n\n")

    def make_items(self, items: List[str]) -> None:
        for item in items:
            self.__sb.append("  ・")
            self.__sb.append(item + "\n")

        self.__sb.append("\n")

    def close(self) -> None:
        self.__sb.append("======================================\n")

    def get_text_result(self) -> str:
        return self.__sb.string
