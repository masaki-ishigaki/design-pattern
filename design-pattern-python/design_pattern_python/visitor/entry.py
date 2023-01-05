from abc import ABCMeta, abstractmethod
from typing import TYPE_CHECKING, Iterator, List

if TYPE_CHECKING:
    from visitor import Visitor


class Element(metaclass=ABCMeta):
    @abstractmethod
    def accept(self, v: "Visitor") -> None:
        pass


class Entry(Element, metaclass=ABCMeta):
    @abstractmethod
    def get_name(self) -> str:
        pass

    @abstractmethod
    def get_size(self) -> int:
        pass

    def to_string(self) -> str:
        return self.get_name() + " (" + str(self.get_size()) + ")"


class File(Entry):
    def __init__(self, name: str, size: int) -> None:
        self.__name = name
        self.__size = size

    def get_name(self) -> str:
        return self.__name

    def get_size(self) -> int:
        return self.__size

    def accept(self, v: "Visitor") -> None:
        v.visit_file(self)


class Directory(Entry):
    def __init__(self, name: str) -> None:
        self.__name = name
        self.__directory: List[Entry] = []

    def get_name(self) -> str:
        return self.__name

    def get_size(self) -> int:
        size = 0
        for entry in self.__directory:
            size += entry.get_size()
        return size

    def add(self, entry: Entry) -> Entry:
        self.__directory.append(entry)
        return self

    def __iter__(self) -> Iterator[Entry]:
        yield from self.__directory

    def accept(self, v: "Visitor") -> None:
        v.visit_directory(self)
