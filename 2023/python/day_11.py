from collections import defaultdict


def expand_universe(grid):
    i = 0
    j = 0
    while True:
        row = grid[j]
        col = [grid_row[i] for grid_row in grid]
        if "#" not in row:
            grid.insert(j, ["." for _ in range(len(row))])
            j += 1
        if "#" not in col:
            for grid_row in grid:
                grid_row.insert(i, ".")
            i += 1
        i += 1
        j += 1
        if i == len(grid) or i >= len(grid[0]):
            break

    gal_coords = defaultdict() 
    c = 1
    for i, line in enumerate(grid):
        for j, dot in enumerate(line):
            if dot == "#":
                line[j] = str(c)
                gal_coords[str(c)] = (i, j)
               
                c += 1

    result = 0
    nums = [str(i) for i in range(1, c + 1)]
    for k, v in gal_coords.items():
        for k2, v2 in gal_coords.items():
            if k == k2:
                continue
            if k2 in nums:
                res = abs(v2[1] - v[1]) + abs(v2[0] - v[0])
                # print(f"{k} - {k2} -> {res}")
                result += res
        nums.remove(k)

    return result

def expand_universe_million(grid):
    i = 0
    j = 0
    cols_mod = []
    rows_mod = []
    while True:
        row = grid[j]
        col = [grid_row[i] for grid_row in grid]
        if "#" not in row:
            rows_mod.append(j)
        if "#" not in col:
            cols_mod.append(i)
        i += 1
        j += 1
        if i == len(grid) or i >= len(grid[0]):
            break

    # grid_list = ["".join(i) for i in grid]
    # for grid_str in grid_list:
    #     print(grid_str)
    gal_coords = defaultdict() 
    c = 1
    for i, line in enumerate(grid):
        for j, dot in enumerate(line):
            if dot == "#":
                line[j] = str(c)
                gal_coords[str(c)] = (i, j)
               
                c += 1

    result = 0
    nums = [str(i) for i in range(1, c + 1)]
    for k, v in gal_coords.items():
        for k2, v2 in gal_coords.items():
            if k == k2:
                continue
            if k2 in nums:
                result += range_mul(v[0], v2[0], rows_mod)
                result += range_mul(v[1], v2[1], cols_mod)
                res = abs(v2[1] - v[1]) + abs(v2[0] - v[0])
                # print(f"{k} - {k2} -> {res}")
                result += res
        nums.remove(k)

    return result

def range_mul(v1, v2, list_existing):
    l = [int(v1), int(v2)]
    ran = range(*sorted(l))

    ran_c = [1_000_000 - 1 for i in list_existing if i in ran] 
    return sum(ran_c)

if __name__ == "__main__":
    with open("data/day_11.txt") as f:

        grid = [list(row.strip()) for row in f]
        num_gal = expand_universe(grid)

        print(num_gal)
        print("=" * 80)

    with open("data/day_11.txt") as f:

        grid = [list(row.strip()) for row in f]
        num_gal = expand_universe_million(grid)

        print(num_gal)


        # grid_list = ["".join(i) for i in grid]
        # for grid_str in grid_list:
        #     print(grid_str)
