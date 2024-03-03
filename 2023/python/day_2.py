from functools import reduce


def power_line(row):
    color_dict = {"blue": 0, "red": 0, "green": 0}

    # Game id
    _, rest = row.split(":")

    # colors
    colors = rest.split(";")

    for game in colors:
        for color in game.split(","):
            num, col = color.strip().split()
            if int(num) > color_dict[col]:
                color_dict[col] = int(num)
    return reduce(lambda x, y: x * y, color_dict.values())


def parse_line(row):
    color_dict = {"blue": 14, "red": 12, "green": 13}

    # Game id
    id_string, rest = row.split(":")

    game_id = int(id_string.split()[-1])

    # colors
    colors = rest.split(";")

    for game in colors:
        for color in game.split(","):
            num, col = color.strip().split()
            if int(num) > color_dict[col]:
                return False

    return game_id


if __name__ == "__main__":
    with open("data/day_2.txt") as f:
        c = 0
        c2 = 0
        for row in f:
            game_id = parse_line(row)
            power_num = power_line(row)
            if game_id:
                c += game_id
            c2 += power_num

        print(c)
        print(c2)
