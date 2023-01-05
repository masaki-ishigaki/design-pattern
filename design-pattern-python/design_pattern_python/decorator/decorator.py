from abc import ABCMeta, abstractmethod


class Display(metaclass=ABCMeta):
    @abstractmethod
    def get_columns(self) -> int:
        pass

    @abstractmethod
    def get_rows(self) -> int:
        pass

    @abstractmethod
    def get_row_text(self, row: int) -> str:
        pass

    def show(self) -> None:
        for i in range(self.get_rows()):
            print(self.get_row_text(i))


class StringDisplay(Display):
    def __init__(self, string: str) -> None:
        self.__string = string

    def get_columns(self) -> int:
        return len(self.__string)

    def get_rows(self) -> int:
        return 1

    def get_row_text(self, row: int) -> str:
        if row != 0:
            raise IndexError()
        return self.__string


class Border(Display):
    def __init__(self, display: Display) -> None:
        self._display = display


class SideBorder(Border):
    def __init__(self, display: Display, ch: str) -> None:
        super().__init__(display)
        self.__border_char = ch

    def get_columns(self) -> int:
        return 1 + self._display.get_columns() + 1

    def get_rows(self) -> int:
        return self._display.get_rows()

    def get_row_text(self, row: int) -> str:
        return self.__border_char + self._display.get_row_text(row) + self.__border_char


class FullBorder(Border):
    def __init__(self, display: Display) -> None:
        super().__init__(display)

    def get_columns(self) -> int:
        return 1 + self._display.get_columns() + 1

    def get_rows(self) -> int:
        return 1 + self._display.get_rows() + 1

    def get_row_text(self, row: int) -> str:
        if row == 0:
            return "+" + self.__make_line("-", self._display.get_columns()) + "+"
        elif row == (self._display.get_rows() + 1):
            return "+" + self.__make_line("-", self._display.get_columns()) + "+"
        else:
            return "|" + self._display.get_row_text(row - 1) + "|"

    def __make_line(self, ch: str, count: int) -> str:
        string = ""
        for _ in range(count):
            string += ch
        return string


if __name__ == "__main__":
    b1 = StringDisplay("Hello, world.")
    b2 = SideBorder(b1, "#")
    b3 = FullBorder(b2)
    b1.show()
    b2.show()
    b3.show()
    b4 = SideBorder(
        FullBorder(
            FullBorder(SideBorder(FullBorder(StringDisplay("Hello, world.")), "*"))
        ),
        "/",
    )
    b4.show()
