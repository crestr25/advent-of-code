import re
from functools import reduce

symbols_str = r"[^a-zA-z0-9_.\n]"

def check_gears(rows):
    res = []
    for i in range(len(rows)):
        items = re.finditer(r"[*]", rows[i])
        for match in items:
            local_res = []
            ind_i = match.start() 
            ind_f = match.end()
            
            if ind_i > 0:
                if bool(re.search(r"[0-9]", rows[i][ind_i - 1])):
                    local_res.append(next(x[0] for x in re.finditer(r"[0-9]+", rows[i]) if x.end() == ind_i))

            if ind_f < len(rows[i]):
                if bool(re.search(r"[0-9]", rows[i][ind_f])):
                    local_res.append(next(x[0] for x in re.finditer(r"[0-9]+", rows[i]) if x.start() == ind_f))

            if i != 0:
                exists = check_line_2(rows[i - 1], ind_i, ind_f)
                if exists:
                    local_res.extend(exists)

            if i < len(rows) - 1:
                exists = check_line_2(rows[i + 1], ind_i, ind_f)
                if exists:
                    local_res.extend(exists)

            if len(local_res) >= 2:
                res.append(reduce(lambda x, y: int(x) * int(y), local_res))
    return res

def check_line_2(s, start, stop):
    if start > 0 and stop < len(s):
        range_2 = range(start - 1 , stop + 1)
    else:
        if start == 0:
            range_2 = range(start , stop + 1)
        elif stop >= len(s):
            range_2 = range(start - 1 , stop)
    print(range_2)
    return [x[0] for x in re.finditer(r"[0-9]+", s) if x.start() in range_2 or x.end() - 1 in range_2]

def check_line(s, start, stop):
    if start > 0 and stop < len(s):
        sub_s = s[start - 1 : stop + 1]
    else:
        if start == 0:
            sub_s = s[start : stop + 1]
        elif stop >= len(s):
            sub_s = s[start - 1 : stop]
    return bool(re.search(symbols_str, sub_s))


def check_valid_parts(rows):
    res = []
    for i in range(len(rows)):
        items = re.findall("[0-9]+", rows[i])
        print(f" {i}===============")
        for item in items:
            match = re.search(f"\\b{item}\\b", rows[i])
            ind_i = match.start()
            if ind_i > 0:
                if bool(re.search(symbols_str, rows[i][ind_i - 1])):
                    res.append(int(item))
                    continue

            ind_f = match.end()
            if ind_f < len(rows[i]):
                if bool(re.search(symbols_str, rows[i][ind_f])):
                    res.append(int(item))
                    continue

            # check line before
            if i != 0:
                exists = check_line(rows[i - 1], ind_i, ind_f)
                if exists:
                    res.append(int(item))
                    continue

            if i < len(rows) - 1:
                exists = check_line(rows[i + 1], ind_i, ind_f)
                if exists:
                    res.append(int(item))
                    continue

    return res


if __name__ == "__main__":
    with open("2023/data/day_3.txt") as f:
        results = []

        file_rows = [row.strip() for row in f]

        res = check_valid_parts(file_rows)

        print(sum(res))

        res = check_gears(file_rows)
        print(res)
        print(sum(res))
