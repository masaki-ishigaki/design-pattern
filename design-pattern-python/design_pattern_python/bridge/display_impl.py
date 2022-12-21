from abc import ABCMeta, abstractmethod


class DisplayImpl(metaclass=ABCMeta):
    @abstractmethod
    def raw_open(self):
        pass

    @abstractmethod
    def raw_print(self):
        pass

    @abstractmethod
    def raw_close(self):
        pass


class StringDisplayImpl(DisplayImpl):
    def __init__(self, string: str) -> None:
        self.__string = string
        self.__width = len(string)

    def raw_open(self) -> None:
        self.print_line()

    def raw_print(self) -> None:
        print("|" + self.__string + "|")

    def raw_close(self) -> None:
        self.print_line()

    def print_line(self):
        print("+", end="")
        for _ in range(self.__width):
            print("-", end="")
        print("+")
