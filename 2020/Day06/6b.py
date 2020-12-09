f = open("6.txt")
group = dict()
votes = set()
totals = []
total = 0
member = 0

for l in f:
    line = l.strip()
    if not line: #if newline, validate the ifno
        x = len(group) - 1
        answer = set()
        if x == 0: answer = group[0]
        elif x < 0: print("!!!!!ERROR!!!!!")
        else:
            answer = group[x]
            while x > 0:
                answer = answer.intersection(group[x-1])
                x += - 1
        totals.append(len(answer))
        total += len(answer)
        group.clear()
        member = 0
    else: #store votes in a set dictonary indexed by member name
        for x in range(len(line)): votes.add(line[x])
        group[member] = votes.copy()
        votes.clear()
        member += 1


print("\n" + str(total) + " questions were answered 'yes' to by everyone.\n")


f.close()
