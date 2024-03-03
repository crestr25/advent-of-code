def move_first(coords, direction, step):
    visited = []
    for _ in range(step):
        if direction == "U":
            coords[0][1] += 1
        elif direction == "D":
            coords[0][1] -= 1
        elif direction == "R":
            coords[0][0] += 1
        elif direction == "L":
            coords[0][0] -= 1

        move_second(coords)
        visited.append(tuple(coords[-1]))

    return visited


def move_second(coords):
    if len(coords) == 1:
        return
    direction = check_adjacent(coords[0], coords[1])
    if direction == "U":
        coords[1][0] = coords[0][0]
        coords[1][1] = coords[0][1] - 1
    elif direction == "D":
        coords[1][0] = coords[0][0]
        coords[1][1] = coords[0][1] + 1
    elif direction == "R":
        coords[1][0] = coords[0][0] - 1
        coords[1][1] = coords[0][1]
    elif direction == "L":
        coords[1][0] = coords[0][0] + 1
        coords[1][1] = coords[0][1]
    elif direction == "UR":
        coords[1][0] = coords[1][0] + 1
        coords[1][1] = coords[1][1] + 1
    elif direction == "UL":
        coords[1][0] = coords[1][0] - 1
        coords[1][1] = coords[1][1] + 1
    elif direction == "DR":
        coords[1][0] = coords[1][0] + 1
        coords[1][1] = coords[1][1] - 1
    elif direction == "DL":
        coords[1][0] = coords[1][0] - 1
        coords[1][1] = coords[1][1] - 1

    move_second(coords[1:])


def check_adjacent(a, b):
    if a[0] == b[0] and (a[1] - b[1]) == 2:
        return "U"
    elif a[0] == b[0] and (a[1] - b[1]) == -2:
        return "D"
    elif a[1] == b[1] and (a[0] - b[0]) == 2:
        return "R"
    elif a[1] == b[1] and (a[0] - b[0]) == -2:
        return "L"

    try:
        if abs(a[1] - b[1]) / abs(a[0] - b[0]) == 2:
            if a[1] > b[1]:
                return "U"

            if a[1] < b[1]:
                return "D"
        elif abs(a[1] - b[1]) / abs(a[0] - b[0]) == 0.5:
            if a[0] > b[0]:
                return "R"

            if a[0] < b[0]:
                return "L"
        elif abs(a[1] - b[1]) == 2 and abs(a[0] - b[0]) == 2:
            if abs(a[1] - b[1]) / abs(a[0] - b[0]) == 1:
                if a[0] > b[0] and a[1] > b[1]:
                    return "UR"
                if a[0] < b[0] and a[1] > b[1]:
                    return "UL"
                if a[0] > b[0] and a[1] < b[1]:
                    return "DR"
                if a[0] < b[0] and a[1] < b[1]:
                    return "DL"
        
    except ZeroDivisionError:
        pass


if __name__ == "__main__":
    grid = [["." for _ in range(30)] for _ in range(30)]

    print("Day 9: Rope Bridge")
    print("=" * 30)
    print("Part 1")

    with open("data/bridge.txt") as file:
        coords = [[0, 0] for _ in range(2)]
        points_visited = []

        for line in file:
            direction, step = line.strip().split()

            visited_t = move_first(coords, direction, int(step))
            points_visited.extend(visited_t)

    print(len(set(points_visited)))

    print("=" * 30)
    print("Part 2")

    with open("data/bridge.txt") as file:
        coords = [[0, 0] for _ in range(10)]
        points_visited = []

        for line in file:
            direction, step = line.strip().split()
            visited_t = move_first(coords, direction, int(step))
            points_visited.extend(visited_t)

    print(len(set(points_visited)))
