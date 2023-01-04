from hand import Hand
from strategy import ProbStrategy, Strategy, WinningStrategy


class Player:
    def __init__(self, name: str, strategy: Strategy) -> None:
        self.__name = name
        self.__strategy = strategy
        self.__win_count = 0
        self.__lose_count = 0
        self.__game_count = 0

    def next_hand(self) -> Hand:
        return self.__strategy.next_hand()

    def win(self):
        self.__strategy.study(True)
        self.__win_count += 1
        self.__game_count += 1

    def lose(self):
        self.__strategy.study(False)
        self.__lose_count += 1
        self.__game_count += 1

    def even(self):
        self.__game_count += 1

    def __str__(self):
        return "[{name}:{game_count} games, {win_count} win, {lose_count} lose]".format(
            name=self.__name,
            game_count=self.__game_count,
            win_count=self.__win_count,
            lose_count=self.__lose_count,
        )


if __name__ == "__main__":
    player1 = Player("Taro", WinningStrategy())
    player2 = Player("Hana", ProbStrategy())

    for i in range(10000):
        next_hand_1 = player1.next_hand()
        next_hand_2 = player2.next_hand()

        if next_hand_1.is_stronger_than(next_hand_2):
            print("Winner:", player1)
            player1.win()
            player2.lose()
        elif next_hand_2.is_stronger_than(next_hand_1):
            print("Winner:", player2)
            player1.lose()
            player2.win()
        else:
            print("Even...")
            player1.even()
            player2.even()

    print("Total result:")
    print(player1)
    print(player2)
