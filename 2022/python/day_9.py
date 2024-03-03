
if __name__ == "__main__":
    print("Day 10: Cathode-Ray Tubes")
    print("=" * 30)
    print("Part 1")
    with open("data/cathode.txt") as file:
        cycles = 1
        X = 1
        active = False
        value = 0
        result = []
        try:
            while True:
                if cycles in (20, 60, 100, 140, 180, 220):
                    result.append(cycles * X)
                if active:
                    X += value
                    active = False
                    cycles += 1
                    continue
                instruction, *args = next(file).strip().split()
                if instruction == "noop":
                    cycles += 1
                    continue
                active = True
                value = int(args[0])
                cycles += 1

        except StopIteration:
            pass
        print(sum(result))

    print("=" * 30)
    print("Part 2")

    with open("data/cathode.txt") as file:
        cycles = 1
        X = 1
        active = False
        value = 0
        result = []
        try:
            while True:
                if cycles in (41, 81, 121, 161, 201, 241):
                    result.append(cycles * X)
                if active:
                    X += value
                    active = False
                    cycles += 1
                    continue
                instruction, *args = next(file).strip().split()
                if instruction == "noop":
                    cycles += 1
                    continue
                active = True
                value = int(args[0])
                cycles += 1

        except StopIteration:
            pass
        print(sum(result))
