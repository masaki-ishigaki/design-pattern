import random
import time
from typing import List


class Memento:
    def __init__(self, money: int) -> None:
        self.__money = money
        self.__fruits: List[str] = []

    def get_money(self) -> int:
        return self.__money

    def add_fruit(self, fruit: str) -> None:
        self.__fruits.append(fruit)

    def get_fruits(self) -> List[str]:
        return self.__fruits


class Gamer:
    fruits_name = ["リンゴ", "ぶどう", "バナナ", "みかん"]

    def __init__(self, money: int) -> None:
        self.__money = money
        self.__fruits: List[str] = []

    def get_money(self) -> int:
        return self.__money

    def bet(self) -> None:
        dice = random.randint(1, 7)
        if dice == 1:
            self.__money += 100
            print("所持金が増えました。")
        elif dice == 2:
            self.__money = int(self.__money / 2)
            print("所持金が半分になりました。")
        elif dice == 6:
            f = self.__get_fruit()
            print(f"フルーツ({f})をもらいました。")
            self.__fruits.append(f)
        else:
            print("何も起こりませんでした。")

    def create_memento(self) -> Memento:
        m = Memento(self.__money)
        for f in self.__fruits:
            if f.startswith("おいしい"):
                m.add_fruit(f)
        return m

    def restore_memento(self, memento: Memento) -> None:
        self.__money = memento.get_money()
        self.__fruits = memento.get_fruits()

    def get_string(self) -> str:
        string = f"[money= {self.__money} + fruits ="
        for f in self.__fruits:
            string += " " + f
        return string

    def __get_fruit(self) -> str:
        f = self.fruits_name[random.randint(0, len(self.fruits_name) - 1)]
        yum_flg = random.randint(0, 2)
        if yum_flg == 0:
            return "おいしい" + f
        else:
            return f


if __name__ == "__main__":
    gamer = Gamer(100)
    memento = gamer.create_memento()

    # ゲームスタート
    for i in range(100):
        print(f"==== {i}")
        print("現状" + gamer.get_string())

        # ゲームを始める
        gamer.bet()

        print(f"所持金は{gamer.get_money()}円になりました。")

        # Mementoの取り扱いの決定
        if gamer.get_money() > memento.get_money():
            print("※だいぶ増えたので、現在の状態を保存しておこう！")
            memento = gamer.create_memento()
        elif gamer.get_money() < (memento.get_money() / 2):
            print("だいぶ減ったので、以前の状態を復元しよう！")
            gamer.restore_memento(memento)
        else:
            pass

        time.sleep(0.1)
