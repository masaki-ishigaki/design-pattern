from abc import ABCMeta, abstractmethod
from typing import List


class Entry(metaclass=ABCMeta):
    @abstractmethod
    def get_name(self) -> str:
        pass

    @abstractmethod
    def get_size(self) -> int:
        pass

    @abstractmethod
    def print_list(self, prefix: str) -> None:
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

    def print_list(self, prefix: str = "") -> None:
        print(prefix + "/" + self.to_string())


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

    def print_list(self, prefix: str = "") -> None:
        print(prefix + "/" + self.to_string())
        for entry in self.__directory:
            entry.print_list(prefix + "/" + self.__name)

    def add(self, entry: Entry) -> Entry:
        self.__directory.append(entry)
        return self


if __name__ == "__main__":
    print("Making root entries...")
    rootdir = Directory("root")
    bindir = Directory("bin")
    tmpdir = Directory("tmp")
    usrdir = Directory("usr")
    rootdir.add(bindir)
    rootdir.add(tmpdir)
    rootdir.add(usrdir)
    bindir.add(File("vi", 10000))
    bindir.add(File("latex", 20000))
    rootdir.print_list()
    print()

    print("Making user entries...")
    yuki = Directory("yuki")
    hanako = Directory("hanako")
    tomura = Directory("tomura")
    usrdir.add(yuki)
    usrdir.add(hanako)
    usrdir.add(tomura)
    yuki.add(File("diary.html", 100))
    yuki.add(File("Composite.java", 200))
    hanako.add(File("memo.tex", 300))
    tomura.add(File("game.doc", 400))
    tomura.add(File("junk.mail", 500))
    rootdir.print_list()
