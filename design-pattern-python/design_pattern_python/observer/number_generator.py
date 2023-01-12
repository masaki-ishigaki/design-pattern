import random
from abc import ABCMeta, abstractmethod
from typing import TYPE_CHECKING, List

if TYPE_CHECKING:
    from observer import Observer


class NumberGenerator(metaclass=ABCMeta):
    def __init__(self) -> None:
        self.__observers: List[Observer] = []

    def add_observer(self, observer: "Observer") -> None:
        self.__observers.append(observer)

    def delete_observer(self, observer: "Observer") -> None:
        self.__observers.remove(observer)

    def notify_observers(self) -> None:
        for o in self.__observers:
            o.update(self)

    @abstractmethod
    def get_number(self) -> int:
        pass

    @abstractmethod
    def execute(self) -> None:
        pass


class RandomNumberGenerator(NumberGenerator):
    def __init__(self):
        super().__init__()
        self.__number = 0

    def get_number(self) -> int:
        return self.__number

    def execute(self) -> None:
        for _ in range(20):
            self.__number = random.randint(0, 50)
            self.notify_observers()
