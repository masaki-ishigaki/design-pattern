from abc import ABCMeta, abstractmethod

from typing_extensions import final


class Product(metaclass=ABCMeta):
    @abstractmethod
    def use(self):
        pass

    @abstractmethod
    def __str__(self):
        pass


class Factory(metaclass=ABCMeta):
    @abstractmethod
    def create_product(self, owner: str) -> Product:
        pass

    @abstractmethod
    def register_product(self, product: Product) -> None:
        pass

    @final
    def create(self, owner: str) -> Product:
        p = self.create_product(owner)
        self.register_product(p)
        return p


class IDCard(Product):
    def __init__(self, owner: str) -> None:
        print(owner + "のカードを作ります。")
        self.__owner = owner

    def use(self) -> None:
        print(self, end="")
        print("を使います。")

    def __str__(self) -> str:
        return "[IDCard:" + self.__owner + "]"

    @property
    def owner(self):
        return self.__owner


class IDCardFactory(Factory):
    def create_product(self, owner: str) -> Product:
        return IDCard(owner)

    def register_product(self, product: Product) -> None:
        print(product, end="")
        print("を登録しました。")


if __name__ == "__main__":
    factory = IDCardFactory()
    card1 = factory.create("Hiroshi Yuki")
    card2 = factory.create("Tomura")
    card3 = factory.create("Hanako Sato")

    card1.use()
    card2.use()
    card3.use()
