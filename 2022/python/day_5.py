if __name__ == "__main__":

    print("Day 5: Tuning Trouble")
    print("=" * 30)
    print("Part 1")
    with open("data/stream.txt") as f:
        stream = f.readline()
        l = (i+4 for i in range(len(stream) - 4) if len(set(stream[i:i+4])) == 4)

    print(next(l))

    print("=" * 30)
    print("Part 2")
    with open("data/stream.txt") as f:
        stream = f.readline()
        l = (i+14 for i in range(len(stream) - 14) if len(set(stream[i:i+14])) == 14)

    print(next(l))
