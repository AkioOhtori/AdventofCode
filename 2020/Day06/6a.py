f = open("6.txt")
group = set()
total = 0

for l in f:
    line = l.strip()
    if not line: #if blank line, validate the data
        #print(group)
        total += len(group)
        group.clear()
    else: 
        for x in range(len(line)): group.add(line[x])

print("\n" + str(total) + " questions were answered 'yes' to.\n")

f.close()
