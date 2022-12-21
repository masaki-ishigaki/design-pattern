from display_impl import DisplayImpl, StringDisplayImpl
from typing_extensions import final


class Display:
    def __init__(self, impl: DisplayImpl) -> None:
        self.__impl = impl

    def open(self) -> None:
        self.__impl.raw_open()

    def print(self) -> None:
        self.__impl.raw_print()

    def close(self) -> None:
        self.__impl.raw_close()

    @final
    def display(self) -> None:
        self.open()
        self.print()
        self.close()


class CountDisplay(Display):
    def multi_display(self, times: int) -> None:
        self.open()
        for _ in range(times):
            self.print()
        self.close()


if __name__ == "__main__":
    d1 = Display(StringDisplayImpl("Hello, Japan."))
    d2 = CountDisplay(StringDisplayImpl("Hello, World."))
    d3 = CountDisplay(StringDisplayImpl("Hello, Universe."))
    d1.display()
    d2.display()
    d3.display()
    d3.multi_display(5)
