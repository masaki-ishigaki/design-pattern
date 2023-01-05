from abc import ABCMeta, abstractmethod
from typing import Any


class Iterator(metaclass=ABCMeta):
    @abstractmethod
    def hasNext(self):
        pass

    @abstractmethod
    def next(self):
        pass


class Aggregate(metaclass=ABCMeta):
    @abstractmethod
    def iterator(self):
        pass


class Book:
    def __init__(self, name: str) -> None:
        self.__name = name

    @property
    def name(self) -> str:
        return self.__name


class BookShelfIterator(Iterator):
    def __init__(self, book_shelf: Any) -> None:
        self.__book_shelf = book_shelf
        self.__index = 0

    def hasNext(self) -> bool:
        return True if self.__index < self.__book_shelf.get_length() else False

    def next(self) -> Book:
        book = self.__book_shelf.get_book_at(self.__index)
        self.__index += 1
        return book


class BookShelf(Aggregate):
    def __init__(self, maxsize: int):
        self.__books = [Book("")] * maxsize
        self.__last = 0

    def get_book_at(self, index: int) -> Book:
        if index < len(self.__books):
            return self.__books[index]
        else:
            raise Exception("Index is out of bookshelf length!!")

    def append_book(self, book: Book) -> None:
        if self.__last < len(self.__books):
            self.__books[self.__last] = book
            self.__last += 1

    def get_length(self) -> int:
        return len(self.__books)

    def iterator(self) -> BookShelfIterator:
        return BookShelfIterator(self)


if __name__ == "__main__":
    book_shelf = BookShelf(4)
    book_shelf.append_book(Book("Around the World in 80 days"))
    book_shelf.append_book(Book("Bible"))
    book_shelf.append_book(Book("Cinderella"))
    book_shelf.append_book(Book("Daddy-Long-Legs"))

    it = book_shelf.iterator()
    while it.hasNext():
        book = it.next()
        print(book.name)
