class Book:
    def __init__(self, name: str):
        self.__name = name

    @property
    def name(self) -> str:
        return self.__name


class BookShelf:
    def __init__(self, maxsize: int):
        self.__books = [Book("")] * maxsize
        self.__last = 0
        self.__index = 0

    def __iter__(self):
        return self

    def __next__(self) -> Book:
        if self.__index == len(self.__books):
            raise StopIteration()

        book = self.__books[self.__index]
        self.__index += 1
        return book

    def get_book_at(self, index: int) -> Book:
        if index < len(self.__books):
            return self.__books[index]
        else:
            raise Exception("Index is out of bookshelf length!!")

    def append_book(self, book: Book) -> None:
        if self.__last < len(self.__books):
            self.__books[self.__last] = book
            self.__last += 1


if __name__ == "__main__":
    book_shelf = BookShelf(4)
    book_shelf.append_book(Book("Around the World in 80 days"))
    book_shelf.append_book(Book("Bible"))
    book_shelf.append_book(Book("Cinderella"))
    book_shelf.append_book(Book("Daddy-Long-Legs"))

    for book in book_shelf:
        print(book.name)
