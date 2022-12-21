from abc import ABCMeta, abstractmethod
from importlib import import_module
from typing import List


class Item(metaclass=ABCMeta):
    def __init__(self, caption) -> None:
        self.caption = caption

    @abstractmethod
    def makeHTML(self) -> str:
        pass


class Link(Item, metaclass=ABCMeta):
    def __init__(self, caption, url: str) -> None:
        self.url = url
        super().__init__(caption)


class Tray(Item, metaclass=ABCMeta):
    def __init__(self, caption: str) -> None:
        self.tray: List[Item] = []
        super().__init__(caption)

    def add(self, item: Item) -> None:
        self.tray.append(item)


class Page(metaclass=ABCMeta):
    def __init__(self, title: str, author: str) -> None:
        self.title = title
        self.author = author
        self.content: List[Item] = []

    def add(self, item: Item) -> None:
        self.content.append(item)

    def output(self, filename: str) -> None:
        with open(filename, "w") as f:
            f.write(self.makeHTML())
        print(filename + "を作成しました。")

    @abstractmethod
    def makeHTML(self) -> str:
        pass


class Factory(metaclass=ABCMeta):
    @classmethod
    def get_factory(cls, classname: str) -> "Factory":
        module, class_ = classname.rsplit(".", 1)
        return getattr(import_module(module), class_)()

    @abstractmethod
    def create_link(self, caption: str, url: str) -> Link:
        pass

    @abstractmethod
    def create_tray(self, caption: str) -> Tray:
        pass

    @abstractmethod
    def create_page(self, title: str, author: str) -> Page:
        pass
