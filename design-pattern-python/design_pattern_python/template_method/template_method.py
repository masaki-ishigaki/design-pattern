from abc import ABCMeta, abstractmethod


class AbstractDisplay(metaclass=ABCMeta):
    @abstractmethod
    def begin(self):
        pass

    @abstractmethod
    def output(self):
        pass

    @abstractmethod
    def end(self):
        pass

    def display(self):
        self.begin()
        for i in range(5):
            self.output()
        self.end()


class CharDisplay(AbstractDisplay):
    def __init__(self, ch: str):
        self.__ch = ch

    def begin(self):
        print("<<", end="")

    def output(self):
        print(self.__ch, end="")

    def end(self):
        print(">>")


class StringDisplay(AbstractDisplay):
    def __init__(self, string: str):
        self.__string = string
        self.__width = len(string)

    def begin(self):
        self.print_line()

    def output(self):
        print("|" + self.__string + "|")

    def end(self):
        self.print_line()

    def print_line(self):
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
