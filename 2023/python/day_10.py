dir_dict = {
    "|": "NS",
    "-": "EW",
    "L": "NE",
    "J": "NW",
    "7": "SW",
    "F": "SE",
}

dir_map = {
    "N": "S",
    "S": "N",
    "E": "W",
    "W": "E",
}

move_dict = {
    "N": lambda x: [x[0] - 1, x[1]],
    "S": lambda x: [x[0] + 1, x[1]],
    "E": lambda x: [x[0], x[1] + 1],
    "W": lambda x: [x[0], x[1] - 1],
}


def posible_starting(coord, grid):
    for dir, func in move_dict.items():
        new_coord = func(coord)
        if -1 in new_coord:
            continue
        pipe = grid[new_coord[0]][new_coord[1]]
        if pipe != ".":
            direction = dir_dict[pipe].replace(dir_map[dir], "")
            if len(direction) == 2:
                continue
            c = move(grid, new_coord, direction)
            if c:
                return c//2



def move(grid, initial_coord, dir):

    coord = move_dict[dir](initial_coord)
    c = 2
    pipe = grid[coord[0]][coord[1]]
    dir = dir_dict[pipe].replace(dir_map[dir], "")

    while True:
        coord = move_dict[dir](coord)
        c += 1
        pipe = grid[coord[0]][coord[1]]
        if pipe == "S":
            break
        dir = dir_dict[pipe].replace(dir_map[dir], "")

    return c

def replace_tiles(grid, initial_coord):

    for dir, func in move_dict.items():
        coord = func(initial_coord)
        if -1 in coord:
            continue
        pipe = grid[coord[0]][coord[1]]
        if pipe != ".":
            dir = dir_dict[pipe].replace(dir_map[dir], "")
            if dir != "S":
                continue
            break

    grid[coord[0]][coord[1]] = "d"
    prev_d = True
    prev_u = False

    while True:
        try:
            coord = move_dict[dir](coord)
            pipe = grid[coord[0]][coord[1]]
            if pipe == "S":
                break
            dir = dir_dict[pipe].replace(dir_map[dir], "")
            if dir == "S":
                grid[coord[0]][coord[1]] = "d"
                prev_d = True
            elif dir == "N":
                grid[coord[0]][coord[1]] = "u"
                prev_u = True
            elif prev_d:
                grid[coord[0]][coord[1]] = "d"
                prev_d = False
            elif prev_u:
                grid[coord[0]][coord[1]] = "u"
                prev_u = False
            else:
                grid[coord[0]][coord[1]] = "*"

                

        except:
            break
    c = 0
    for line in grid:
        for i, dot in enumerate(line):
            if dot not in ("u", "d", "*", "S"):
                for next_dot in range(i + 1, len(line)):
                    if line[next_dot] == "d":
                        line[i] = "I"
                        c += 1
                        break
                    elif line[next_dot] == "u":
                        line[i] = "O"
                        break
                else:
                    line[i] = "O"
                
    grid_list = ["".join(i) for i in grid] 
    for grid_str in grid_list:
        print(grid_str)
    return c

if __name__ == "__main__":
    with open("data/day_10.txt") as f:
        grid = []

        starting_position = []
        for i, row in enumerate(f):
            if "S" in row:
                starting_position.append(i)
                starting_position.append(row.index("S"))
            grid.append(list(row.strip()))

        result = posible_starting(starting_position, grid)
        print(result)
        print('=' * 80)

    with open("data/day_10.txt") as f:
        grid = []

        starting_position = []
        for i, row in enumerate(f):
            if "S" in row:
                starting_position.append(i)
                starting_position.append(row.index("S"))
            grid.append(list(row.strip()))

        # move(grid, starting_position, "S")
        result = replace_tiles(grid, starting_position)

        print(result)
