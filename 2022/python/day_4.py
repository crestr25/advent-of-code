from collections import defaultdict
import re


if __name__ == "__main__":
    crane = defaultdict(list)

    print("Day 4: Supply Stacks")
    print("=" * 30)
    print("Part 1")

    with open("data/crates.txt", "r") as file:
        for line in file:
            if re.sub('[ \n]', '', line).isdigit():  
                next(file)
                break
            for i, box in enumerate(line[1::4], start=1):
                if box != " ":
                    crane[i].insert(0, box)

        for line in file:
            amount, start, stop = re.findall("[0-9]+", line)

            for i in range(int(amount)):
                crane[int(stop)].append(crane[int(start)].pop())
        
        print("".join(crane[i][-1] for i in sorted(crane.keys())))

    print("=" * 30)
    print("Part 2")

    with open("data/crates.txt", "r") as file:
        crane = defaultdict(list)
        for line in file:
            if re.sub('[ \n]', '', line).isdigit():  
                next(file)
                break
            for i, box in enumerate(line[1::4], start=1):
                if box != " ":
                    crane[i].insert(0, box)
        print("=============") 
        print(crane)
        for line in file:
            amount, start, stop = re.findall("[0-9]+", line)
            print(amount, start, stop)
            crane[int(stop)].extend(crane[int(start)][len(crane[int(start)]) - int(amount):])
            del crane[int(start)][len(crane[int(start)]) - int(amount):]
            print(crane)


        print("".join(crane[i][-1] for i in sorted(crane.keys())))
