from collections import defaultdict

def compute_lcm(x, y):

   # choose the greater number
   if x > y:
       greater = x
   else:
       greater = y

   while(True):
       if((greater % x == 0) and (greater % y == 0)):
           lcm = greater
           break
       greater += 1

   return lcm

if __name__ == "__main__":

    # with open("data/day_8.txt") as f:
    #
    #     pattern = next(f).strip()
    #     directions = defaultdict(list)
    #     next(f)
    #
    #     for dir in f:
    #         node, next_dir = dir.split("=")
    #         next_node_dir_l, next_node_dir_r = next_dir[next_dir.find("(")+1:next_dir.find(")")].split(",")
    #         directions[node.strip()].extend([next_node_dir_l.strip(), next_node_dir_r.strip()])
    #
    #     i = 0
    #     location = "AAA"
    #     print(directions)
    #     while True:
    #
    #         new_dir = pattern[i % len(pattern)]
    #         new_node = directions[location]
    #
    #         i += 1
    #
    #         if new_dir == "L":
    #             location = new_node[0]
    #         elif new_dir == "R":
    #             location = new_node[1]
    #
    #         if location == "ZZZ":
    #             break
    #
    #     print(i)
    #     print("=" * 80)

    with open("data/day_8.txt") as f:

        pattern = next(f).strip()
        directions = defaultdict(list)
        next(f)

        for dir in f:
            node, next_dir = dir.split("=")
            next_node_dir_l, next_node_dir_r = next_dir[next_dir.find("(")+1:next_dir.find(")")].split(",")
            directions[node.strip()].extend([next_node_dir_l.strip(), next_node_dir_r.strip()])
        
        locations = [k for k in directions if k[-1] == "A"]
        locations_i = []
        for node in locations:
            i = 0
            while True:
                new_dir_i = pattern[i % len(pattern)]

                new_node_i = directions[node]

                if new_dir_i == "L":
                    node = new_node_i[0]
                elif new_dir_i == "R":
                    node = new_node_i[1]

                i += 1

                if node.endswith("Z"):
                    locations_i.append(i)
                    break
        
        from math import gcd
        lcm = 1
        for i in locations_i:
            lcm = lcm*i//gcd(lcm, i)

        print(lcm)

