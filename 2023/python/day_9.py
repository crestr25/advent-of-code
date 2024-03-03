if __name__ == "__main__":

    with open("data/day_9.txt") as f:
        result = 0
        result_2 = 0
        for row in f:
            nums = list(map(int, row.strip().split()))

            res_list = [nums.copy()]

            while True:
                nums = [nums[i + 1] - nums[i] for i in range(len(nums) - 1)]
                res_list.insert(0, nums)
                if sum([bool(x) for x in nums]) == 0:
                    break
            
            c = res_list[0][0]
            c_1 = res_list[0][0]
            for li in res_list[1:]:
                c = li[-1] + c
                c_1 = li[0] - c_1
            result += c
            result_2 += c_1

        print(result)
        print("="*80)
        print(result_2)
