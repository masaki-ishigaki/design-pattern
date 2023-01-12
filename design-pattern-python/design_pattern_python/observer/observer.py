import time
from abc import ABCMeta, abstractmethod

from number_generator import NumberGenerator, RandomNumberGenerator


class Observer(metaclass=ABCMeta):
    @abstractmethod
    def update(self, generator: NumberGenerator) -> None:
        pass


class DigitObserver(Observer):
    def update(self, generator: NumberGenerator) -> None:
        print(f"DigitObserver:{generator.get_number()}")
        time.sleep(0.1)


class GraphObserver(Observer):
    def update(self, generator: NumberGenerator) -> None:
        print("GraphObserver", end="")
        count = generator.get_number()
        for _ in range(count):
            print("*", end="")
        print("")
        time.sleep(0.1)


if __name__ == "__main__":
    generator = RandomNumberGenerator()
    observer1 = DigitObserver()
    observer2 = GraphObserver()
    generator.add_observer(observer1)
    generator.add_observer(observer2)
    generator.execute()
