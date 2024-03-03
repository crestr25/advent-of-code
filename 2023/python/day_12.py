import re


if __name__ == "__main__":
    with open("data/day_12.txt") as f:

        for line in f:
            springs, groups = line.strip().split()

            incog_spr = re.findall(r"\?+", springs)


            for i in range(len(incog_spr)):
                for spr in incog_spr:
                    x = len(spr)
                    y = x - 1


                    for j in range(x * y):
                        spr = spr.replace("?", "." * x)
                        print(spr)

                        print("=" * 10)


            
