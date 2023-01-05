from abc import ABCMeta, abstractmethod


class AbstractDisplay(metaclass=ABCMeta):
    @abstractmethod
    def begin(self) -> None:
        pass

    @abstractmethod
    def output(self) -> None:
        pass

    @abstractmethod
    def end(self) -> None:
        pass

    def display(self) -> None:
        self.begin()
        for i in range(5):
            self.output()
        self.end()


class CharDisplay(AbstractDisplay):
    def __init__(self, ch: str) -> None:
        self.__ch = ch

    def begin(self) -> None:
        print("<<", end="")

    def output(self) -> None:
        print(self.__ch, end="")

    def end(self) -> None:
        print(">>")


class StringDisplay(AbstractDisplay):
    def __init__(self, string: str) -> None:
        self.__string = string
        self.__width = len(string)

    def begin(self) -> None:
        self.print_line()

    def output(self) -> None:
        print("|" + self.__string + "|")

    def end(self) -> None:
        self.print_line()

    def print_line(self) -> None:
        print("+", end="")
        for _ in range(self.__width):
            print("-", end="")
        print("+")


if __name__ == "__main__":
    d1 = CharDisplay("H")
    d2 = StringDisplay("Hello, world.")

    d1.display()
    print("")
    d2.display()
