
if __name__ == "__main__":
    with open("data/day_4.txt") as f:
        
        res = []
        for row in f:
            c = -1
            _, lines = row.split(":")
            line1, line2 = lines.strip().split("|")
            list_1 = line1.split()
            list_2 = line2.split()

            for i in list_1:
                if i in list_2:
                    c += 1
            if c > -1:
                res.append(2 ** c)

        print(sum(res))

    with open("data/day_4.txt") as f:
        
        res = []
        file_rows = [row.strip() for row in f]

        res = []


        for row in file_rows:

            new_rows = [row]
            ind = 0
            while True:
                
                c = 1 
                try:
                    id_str, lines = new_rows[ind].split(":")
                    id = int(id_str.split()[-1]) - 1
                except Exception:
                    break
                line1, line2 = lines.strip().split("|")
                list_1 = line1.split()
                list_2 = line2.split()
                # print(new_rows)
                for i in list_1:
                    if i in list_2:
                        new_rows.append(file_rows[id + c])
                        c += 1
                    
                ind += 1

            res.extend(new_rows)


        print(len(res))





