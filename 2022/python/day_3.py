def new_range(start, end):
    return range(start, end + 1)

def check_set(a, b):
    return set(a).issubset(set(b)) if len(a) < len(b) else set(b).issubset(set(a))

if __name__ == "__main__":
    # min/max option
    print("Day 3: Camp Clean-Up")
    print("=" * 30)
    print("Part 1")

    data = [line.strip().split(',') for line in open("data/pairs.txt")]
    res = [check_set(new_range(*map(int, b.split('-'))), new_range(*map(int, a.split('-')))) for a, b in data]

    print(sum(res))

    print("=" * 30)
    print("Part 2")

    data = [line.strip().split(',') for line in open("data/pairs.txt")]
    res = [any(set(new_range(*map(int, a.split('-')))) & set(new_range(*map(int, b.split('-'))))) for a, b in data]

    print(sum(res))
