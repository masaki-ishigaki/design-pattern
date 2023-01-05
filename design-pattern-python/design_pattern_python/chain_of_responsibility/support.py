from abc import ABCMeta, abstractmethod
from typing import Optional


class Trouble:
    def __init__(self, number: int) -> None:
        self.__number = number

    def get_number(self) -> int:
        return self.__number

    def to_string(self) -> str:
        return f"[Trouble {self.__number}]"


class Support(metaclass=ABCMeta):
    def __init__(self, name: str) -> None:
        self.__name = name
        self.__next: Optional["Support"] = None

    def set_next(self, next_: "Support") -> "Support":
        self.__next = next_
        return next_

    def support(self, trouble: Trouble) -> None:
        if self.resolve(trouble):
            self.done(trouble)
        elif self.__next is not None:
            self.__next.support(trouble)
        else:
            self.fail(trouble)

    def to_string(self) -> str:
        return f"[{self.__name}]"

    @abstractmethod
    def resolve(self, trouble: Trouble) -> bool:
        pass

    def done(self, trouble: Trouble) -> None:
        print(trouble.to_string() + " is resolved by " + self.to_string() + ".")

    def fail(self, trouble: Trouble) -> None:
        print(trouble.to_string() + " cannot be resolved.")


class NoSupport(Support):
    def __init__(self, name: str) -> None:
        super().__init__(name)

    def resolve(self, trouble: Trouble) -> bool:
        return False


class LimitSupport(Support):
    def __init__(self, name: str, limit: int) -> None:
        super().__init__(name)
        self.__limit = limit

    def resolve(self, trouble: Trouble) -> bool:
        if trouble.get_number() < self.__limit:
            return True
        else:
            return False


class OddSupport(Support):
    def __init__(self, name: str) -> None:
        super().__init__(name)

    def resolve(self, trouble: Trouble) -> bool:
        if trouble.get_number() % 2 == 1:
            return True
        else:
            return False


class SpecialSupport(Support):
    def __init__(self, name: str, number: int) -> None:
        super().__init__(name)
        self.__number = number

    def resolve(self, trouble: Trouble) -> bool:
        if trouble.get_number() == self.__number:
            return True
        else:
            return False


if __name__ == "__main__":
    alice = NoSupport("Alice")
    bob = LimitSupport("Bob", 100)
    charlie = SpecialSupport("Charlie", 429)
    diana = LimitSupport("Diana", 200)
    elmo = OddSupport("Elmo")
    fred = LimitSupport("Fred", 300)

    # 連鎖の作成
    alice.set_next(bob).set_next(charlie).set_next(diana).set_next(elmo).set_next(fred)

    for i in range(500):
        alice.support(Trouble(i))
