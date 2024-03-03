import math
import pandas as pd

def check_grid(grid, i, j):
    column_len = len(grid)
    row_len = len(grid[0])

    scenic_score = []
    blocked = []
    c = 0

    for x_1 in reversed(range(j)):
        c += 1
        if grid[i][j] <= grid[i][x_1]:
            blocked.append(True)
            break

    scenic_score.append(c)
    c = 0

    for x_2 in range(j + 1, row_len):
        c += 1
        if grid[i][j] <= grid[i][x_2]:
            blocked.append(True)
            break

    scenic_score.append(c)
    c = 0

    for y_1 in reversed(range(i)):
        c += 1
        if grid[i][j] <= grid[y_1][j]:
            blocked.append(True)
            break

    scenic_score.append(c)
    c = 0

    for y_2 in range(i + 1, column_len):
        c += 1
        if grid[i][j] <= grid[y_2][j]:
            blocked.append(True)
            break

    scenic_score.append(c)

    return math.prod(scenic_score) if sum(blocked) < 4 else 0


if __name__ == "__main__":
    print("Day 8: Treetop Tree House")
    print("=" * 30)
    print("Part 1")
    data = pd.read_csv("data/trees.txt", header=None)
    print(data)
    with open("data/trees.txt") as f:
        grid = [tuple(map(int, line.strip())) for line in f]

        result = []
        best_tree = 0

        result.extend(grid[0])

        for i in range(1, len(grid) - 1):
            result.append(grid[i][0])
            result.append(grid[i][-1])
            for j in range(1, len(grid[i]) - 1):
                if tree_ss := check_grid(grid, i, j):
                    result.append(grid[i][j])
                    if tree_ss > best_tree:
                        best_tree = tree_ss

        result.extend(grid[-1])
        print(len(result))

    print("=" * 30)
    print("Part 2")

    print(best_tree)
