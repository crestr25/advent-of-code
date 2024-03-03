from functools import reduce


if __name__ == "__main__":
    with open("data/day_6.txt") as f:

        times = [int(i.strip()) for i in next(f).split(":")[-1].strip().split()]
        distance = [int(i.strip()) for i in next(f).split(":")[-1].strip().split()]
        
        res = []
        for i, time in enumerate(times):
            c = 0
            for vel in range(1, time):
                distance_traveled = vel * (time - vel) 
                if distance_traveled > distance[i]:
                    c += 1
            if c:
                res.append(c)

        print(reduce(lambda x, y: x*y, res))

        times = int("".join(list(map(str, times))))
        distance = int("".join(list(map(str, distance))))

        
        res = []
        c = 0
        for vel in range(1, times):
            distance_traveled = vel * (times - vel) 
            if distance_traveled > distance:
                c += 1
        if c:
            res.append(c)

        print(reduce(lambda x, y: x*y, res))
