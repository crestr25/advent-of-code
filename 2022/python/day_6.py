if __name__ == "__main__":

    print("Day 6: No Space Left on Device")
    print("=" * 30)
    print("Part 1")

    with open("data/file_system.txt") as f:
        next(f)
        next(f)
        lines = f.readlines()
        level = ["/"]
        fs_name = "/"
        size = []
        fs = {"/": 0}
        for line in lines:
            line = line.strip().split()
            if line[1] == "cd":
                if line[2] != "..":
                    level.append(f"{level[-1]}-{line[2]}")
                    fs[level[-1]] = 0
                else:
                    level.pop()

            elif line[0].isnumeric():

                for i in level:
                    fs[i] += int(line[0])

        filtered_fs = {}
        for k, v in fs.items():
            if 1 < v <= 100000:
                filtered_fs[k] = v

        print(sum(filtered_fs.values()))

    print("=" * 30)
    print("Part 2")
    system_size = 70000000
    update_size = 30000000

    used_size = system_size - fs["/"] 
    required_size = update_size - used_size

    selected_folder = []
    for k, v in fs.items():
        if v >= required_size and k != "/":
            selected_folder.append((k, v))

    folder_size = min(selected_folder, key=lambda x: x[1])
    print(folder_size)
