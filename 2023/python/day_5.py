import re

if __name__ == "__main__":
    with open("data/day_5.txt") as f:

        _, seeds_str = next(f).split(":")
        seeds = seeds_str.strip().split()
        seeds_range = []
        for i in range(0, len(seeds), 2):
            range_seed = range(int(seeds[i]), int(seeds[i]) + int(seeds[i + 1]))
            seeds_range.append(range_seed)

        print(seeds_range)

        maps = [[] for _ in range(7)]
        i = 0
        next(f)
        while True:
            try:
                row = next(f)
                if row != "\n":
                    if re.match("[0-9]", row):
                        row_nums = row.strip().split()
                        range_1 = range(int(row_nums[1]), int(row_nums[1]) + int(row_nums[2]))
                        range_2 = range(int(row_nums[0]), int(row_nums[0]) + int(row_nums[2]))
                        maps[i].append((range_1, range_2))
                    
                else:
                    i += 1
            except StopIteration:
                break


        res = []
        for seed in seeds:
            new_val = int(seed) 
            for trans in maps:
                for t in trans:
                    if new_val in t[0]:

                        ind = t[0].index(new_val)
                        # print(ind)
                        new_val = t[1][ind]
                        break
            else:
                res.append(new_val)
        print(min(res))

        res = []         
        for seed in range(10000000000000000):
            new_val = seed
            for trans in reversed(maps):
                for t in trans:
                    if new_val in t[1]:

                        ind = t[1].index(new_val)
                        # print(ind)
                        new_val = t[0][ind]
                        break
            else:
                for range_seed in seeds_range:
                    if new_val in range_seed:
                        print(seed, new_val)
                        exit()
        print(new_val)
        # print(min(res))
        # print(res)
        # print(min(res))



