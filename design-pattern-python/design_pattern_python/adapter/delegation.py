from abc import ABCMeta, abstractmethod


class Print(metaclass=ABCMeta):
    @abstractmethod
    def print_weak(self) -> None:
        pass

    @abstractmethod
    def print_strong(self) -> None:
        pass


class Banner:
    def __init__(self, string: str) -> None:
        self.__string = string

    def show_with_paren(self) -> None:
        print("(" + self.__string + ")")

    def show_with_aster(self) -> None:
        print("*" + self.__string + "*")


class PrintBanner(Print):
    def __init__(self, banner: Banner):
        self.__banner = banner

    def print_weak(self) -> None:
        self.__banner.show_with_paren()

    def print_strong(self) -> None:
        self.__banner.show_with_aster()


if __name__ == "__main__":
    p = PrintBanner(Banner("Hello"))
    p.print_weak()
    p.print_strong()
