from enum import Enum


class HandKind(Enum):
    ROCK = ("グー", 0)
    SCISSORS = ("チョキ", 1)
    PAPER = ("パー", 2)

    def __init__(self, ja: str, id: int) -> None:
        self.ja = ja
        self.id = id

    @classmethod
    def get_hand_kind_by_id(cls, id: int) -> "HandKind":
        if id == 0:
            return HandKind.ROCK
        elif id == 1:
            return HandKind.SCISSORS
        elif id == 2:
            return HandKind.PAPER
        else:
            raise Exception("input id is out of range!!")


class Hand:
    def __init__(self, kind: HandKind) -> None:
        self.kind = kind

    def is_stronger_than(self, opponent: "Hand") -> bool:
        return self.__fight(opponent) == 1

    def is_weaker_than(self, opponent: "Hand") -> bool:
        return self.__fight(opponent) == -1

    def __fight(self, opponent: "Hand") -> int:
        if self.kind == opponent.kind:
            return 0
        elif (self.kind.id + 1) % 3 == opponent.kind.id:
            return 1
        else:
            return -1
