from enum import Enum


class Hand(Enum):
    A = 1
    B = 2
    C = 3
    X = 1
    Y = 2
    Z = 3

    @staticmethod
    def match(home, away):
        if (home + 1) % 3 == away % 3:
            value = 6
        elif home == away:
            value = 3
        else:
            value = 0

        return value


class Hand2(Enum):
    A = 1
    B = 2
    C = 3

    @staticmethod
    def match(value, outcome):
        if outcome == "X":
            enum_values = list(Hand2)
            current_index = enum_values.index(value)
            previous_index = (current_index - 1) % len(enum_values)
            res = 0 + enum_values[previous_index].value
        elif outcome == "Y":
            res = 3 + value.value
        else:
            res = 6 + (value.value % 3) + 1
        return res


if __name__ == "__main__":
    print("Day 1: Rock Paper Scissors")
    print("=" * 30)
    print("Part 1")
    with open("2022/data/rock_paper_scissors.txt", "r") as f:
        data = f.read().split("\n")[:-1]

        result = 0
        for pair in data:
            home = Hand[pair.split(" ")[0]].value
            away = Hand[pair.split(" ")[1]].value

            result += Hand.match(home, away) + away
        print(result)

    print("=" * 30)
    print("Part 2")

    with open("2022/data/rock_paper_scissors.txt", "r") as f:
        data = f.read().split("\n")[:-1]

        result = 0
        for pair in data:
            value = Hand2[pair.split(" ")[0]]
            outcome = pair.split(" ")[1]

            result += Hand2.match(value, outcome)

        print(result)
