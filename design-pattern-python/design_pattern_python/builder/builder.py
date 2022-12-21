from abc import ABCMeta, abstractmethod
from typing import List


class Builder(metaclass=ABCMeta):
    @abstractmethod
    def make_title(self, title: str) -> None:
        pass

    @abstractmethod
    def make_string(self, string: str) -> None:
        pass

    @abstractmethod
    def make_items(self, items: List[str]) -> None:
        pass

    @abstractmethod
    def close(self) -> None:
        pass


class StringBuilder:
    def __init__(self):
        self.__string = ""

    def append(self, string: str) -> None:
        self.__string += string

    @property
    def string(self) -> str:
        return self.__string
