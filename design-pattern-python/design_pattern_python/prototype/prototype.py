import copy
from abc import ABCMeta, abstractmethod
from typing import Dict


class Product(metaclass=ABCMeta):
    @abstractmethod
    def use(self, s: str):
        pass

    @abstractmethod
    def createCopy(self):
        pass


class Manager:
    def __init__(self) -> None:
        self.__showcase: Dict[str, Product] = {}

    def register(self, name: str, prototype: Product) -> None:
        self.__showcase[name] = prototype

    def create(self, prototype_name: str) -> Product:
        try:
            p = self.__showcase[prototype_name]
        except KeyError:
            exit()
        return p.createCopy()


class MessageBox(Product):
    def __init__(self, decochar: str) -> None:
        self.__decochar = decochar

    def use(self, s: str):
        decolen = 1 + len(s) + 1
        for _ in range(decolen):
            print(self.__decochar, end="")
        print()
        print(self.__decochar + s + self.__decochar)
        for _ in range(decolen):
            print(self.__decochar, end="")
        print()

    def createCopy(self) -> Product:
        p = copy.deepcopy(self)
        return p


class UnderlinePen(Product):
    def __init__(self, ulchar: str) -> None:
        self.__ulchar = ulchar

    def use(self, s: str):
        ulen = len(s)
        print(s)
        for _ in range(ulen):
            print(self.__ulchar, end="")
        print()

    def createCopy(self):
        p = copy.deepcopy(self)
        return p


if __name__ == "__main__":
    manager = Manager()
    upen = UnderlinePen("-")
    mbox = MessageBox("*")
    sbox = MessageBox("/")

    manager.register("strong message", upen)
    manager.register("warning box", mbox)
    manager.register("slash box", sbox)

    p1 = manager.create("strong message")
    p1.use("Hello, world.")

    p2 = manager.create("warning box")
    p2.use("Hello, world.")

    p3 = manager.create("slash box")
    p3.use("Hello, world.")
