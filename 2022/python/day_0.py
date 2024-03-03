class CalorieListIter:
    def __init__(self, file) -> None:
        self.file = file

    def __iter__(self):
        return self 

    def __next__(self):
        c = 0
        for line in self.file:
            if line.replace("\n", "") == "":
                return c 
            else:
                c += int(line.replace("\n", ""))
        else:
            raise StopIteration


class CalorieList:

    def __init__(self, file_name, mode) -> None:
        self.filename = file_name
        self.mode = mode
        self.file = None
         
    def __enter__(self):
        self.file = open(self.filename, self.mode)
        return CalorieListIter(self.file)
     
    def __exit__(self, exc_type, exc_value, exc_traceback):
        self.file.close()


def easy_solution():

    result = [sum(map(int, n)) for n in [i.split("\n") for i in open("data/info.txt", "r").read()[:-1].split("\n\n")]]
    return result
 

if __name__ == "__main__":
    file_name = "data/info.txt"
    with CalorieList(file_name, "r") as calorie_list:

        max_nums = [int(c) for c in calorie_list]

        max_nums = sorted(max_nums, reverse=True)

        print(f"Max calories carried by an elf: {max_nums[0]}") 
        print(f"Max calories carried by top 3 elf: {sum(max_nums[0:3])}") 

    result = easy_solution()
    result.sort(reverse=True)
    print(f"Quick Solution: {result[0]}")
    print(f"Quick Solution: {sum(result[:3])}")


