
from collections import defaultdict

ind = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"]
ind_joker = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]


def sort_key(string, ind):
    char_to_index = {c: i for i, c in enumerate(ind)}
    indices = [char_to_index.get(c, len(char_to_index)) for c in string]
    return indices

def get_order(hands: list, ind: list):
    sorted_strings = sorted(hands, key=lambda string: sort_key(string[0], ind))
    return sorted_strings

def get_duplicates(s: str) -> dict:
    res = defaultdict(int)

    for char in s:
        res[char] += 1

    return res

def get_hand_type(hand: str):

    dups = get_duplicates(hand)
    if len(dups) == 1:
        return "voak"
    elif len(dups) == 2:
        if 1 not in list(dups.values()):
            return "fh"
        else:
            return "foak"
    elif len(dups) == 3:
        if 3 in dups.values():
            return "tok"
        elif 2 not in dups.values():
            return "hc"
        else:
            return "tp"
    elif len(dups) == 4:
        return "op"
    else:
        return "hc"

def get_duplicates_jok(s: str) -> dict:

    res = defaultdict(int)
    for char in s:
        if char != "J":
            res[char] += 1

    if "J" not in s:
        return res

    if not res:
        return {"J": 5}


    max_key = max(res, key=res.get)
    s = s.replace("J", max_key)

    res = defaultdict(int)
    for char in s:
        res[char] += 1

    return res

def get_hand_type_joker(hand: str):

    dups = get_duplicates_jok(hand)

    if len(dups) == 1:
        return "voak"
    elif len(dups) == 2:
        if 1 not in list(dups.values()):
            return "fh"
        else:
            return "foak"
    elif len(dups) == 3:
        if 3 in dups.values():
            return "tok"
        elif 2 not in dups.values():
            return "hc"
        else:
            return "tp"
    elif len(dups) == 4:
        return "op"
    else:
        return "hc"



if __name__ == "__main__":
    with open("data/day_7.txt") as f:

        dict_map = {
                "voak": [],
                "foak": [], 
                "fh": [],
                "tok": [],
                "tp":[],
                "op": [],
                "hc": []
                }

        total = 0
        for row in f:
            hand, bid = row.split()
            result = get_hand_type(hand)
            dict_map[result].append([hand, bid])
            total += 1

        c = 0
        for hand_type, hands in dict_map.items():
            if hands:
                for hand in get_order(hands, ind):
                    c += (int(hand[1]) * total)
                    total -= 1
        print(c)
        print("=" * 80)

    with open("data/day_7.txt") as f:

        dict_map = {
                "voak": [],
                "foak": [], 
                "fh": [],
                "tok": [],
                "tp":[],
                "op": [],
                "hc": []
                }
        total = 0
        for row in f:
            hand, bid = row.split()
            result = get_hand_type_joker(hand)
            dict_map[result].append([hand, bid])
            total += 1
        
        c = 0
        for hand_type, hands in dict_map.items():
            if hands:
                # print(get_order(hands, ind_joker))
                for hand in get_order(hands, ind_joker):
                    c += (int(hand[1]) * total)
                    total -= 1

        print(c)


                




