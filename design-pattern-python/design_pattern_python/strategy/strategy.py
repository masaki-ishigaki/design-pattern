import random
from abc import ABCMeta, abstractmethod

from hand import Hand, HandKind


class Strategy(metaclass=ABCMeta):
    @abstractmethod
    def next_hand(self) -> Hand:
        pass

    @abstractmethod
    def study(self, win: bool) -> None:
        pass


class WinningStrategy(Strategy):
    def __init__(self) -> None:
        self.__won = False
        self.__prev_hand = Hand(HandKind.ROCK)

    def next_hand(self) -> Hand:
        if not self.__won:
            self.__prev_hand = Hand(HandKind.get_hand_kind_by_id(random.randint(0, 2)))
        return self.__prev_hand

    def study(self, win: bool) -> None:
        self.__won = win


class ProbStrategy(Strategy):
    def __init__(self):
        self.__prev_hand_value = HandKind.ROCK.id
        self.__current_hand_value = HandKind.ROCK.id
        # self.__history[0] -> ROCK
        # self.__history[1] -> SCISSORS
        # self.__history[2] -> PAPER
        self.__history = [
            [1, 1, 1],
            [1, 1, 1],
            [1, 1, 1],
        ]

    def next_hand(self) -> Hand:
        bet = random.randint(0, self.__get_sum(self.__current_hand_value))
        hand_value = HandKind.ROCK.id
        if bet < self.__history[self.__current_hand_value][0]:
            hand_value = HandKind.ROCK.id
        elif (
            bet
            < self.__history[self.__current_hand_value][0]
            + self.__history[self.__current_hand_value][1]
        ):
            hand_value = HandKind.SCISSORS.id
        else:
            hand_value = HandKind.PAPER.id

        self.__prev_hand_value = self.__current_hand_value
        self.__current_hand_value = hand_value

        return Hand(HandKind.get_hand_kind_by_id(hand_value))

    def study(self, win: bool):
        if win:
            self.__history[self.__prev_hand_value][self.__current_hand_value] += 1
        else:
            self.__history[self.__prev_hand_value][
                (self.__current_hand_value + 1) % 3
            ] += 1
            self.__history[self.__prev_hand_value][
                (self.__current_hand_value + 2) % 3
            ] += 1

    def __get_sum(self, id: int) -> int:
        sum = 0
        for i in range(3):
            sum += self.__history[id][i]
        return sum
