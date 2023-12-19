answer = 0
for line in open("Day01\\data.txt"):
    n = []
    for l in line:
        try:
            n.append(int(l))
        except: pass
    answer += int(str(n[0]) + str(n[len(n)-1]))
    print(answer)
    