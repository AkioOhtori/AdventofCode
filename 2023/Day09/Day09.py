tot = 0
PART2 = True

for line in open("Day09\\data.txt"):
    x = 0
    book = {}
    book[x] = list(map(int,line.strip().split()))


    while 1: #calculate until all zeros
        new = []
        # print(book[x])
        for y in range(len(book[x])-1):
            n = book[x][y+1] - book[x][y]
            new.append(n)
        x += 1
        book[x] = new.copy()
        if all(v == 0 for v in new):
            break

    book[x].append(0)
    if not PART2:
        while x != 0:
            x += -1
            n = book[x+1][-1] + book[x][-1]
            book[x].append(n)
        tot += book[0][-1]
    else:    
        while x != 0:
            # print(book[x])
            x += -1
            n = book[x][0] - book[x+1][0]
            book[x].insert(0,n)
        tot += book[0][0]    
print(tot)