# Description: Day 2 of the Advent of Code 2020
import string

mapping = {v: i for i, v in enumerate(string.ascii_letters, start=1)}

def find_similar(items):
   
    res = []
    for item in items:
        items_1 = item[:len(item)//2]
        items_2 = item[len(item)//2:]

        for it in items_1:
            if it in items_2:
                res.append(mapping[it])
                break

    return res

def find_similar2(items):
   
    res = []
    for group in items:
        min_list = list(map(len, group))
        i = min_list.index(min(min_list))

        for item in group[i]:
            if item in group[(i + 1) % 3] and item in group[(i + 2) % 3]:
                res.append(mapping[item])
                break

    return res

if __name__ == "__main__":
    
    print("Day 2: Rucksack Reorganization")
    print("=" * 30)
    print("Part 1")

    with open("data/container.txt") as f:

        items = [line.strip() for line in f.readlines()]
        res = find_similar(items)

        print(sum(res))

    print("=" * 30)
    print("Part 2")

    with open("data/container.txt") as f:

        items = (line.strip() for line in f.readlines())
        by3_items = list(zip(items, items, items))
        res = find_similar2(by3_items)

        print(sum(res))



