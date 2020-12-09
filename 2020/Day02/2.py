f = open("2.txt")
good = 0
bad = 0

test = f.read()
f.close()
full = test.splitlines()

"""
for x in range(len(full)):
    hy = full[x].find("-")
    co = full[x].find(":")

    minn = full[x][0:hy]
    letter = full[x][co-1]
    maxx = full[x][(hy+1):(co-2)]
    pw = full[x][(co+2):]

    num = pw.count(letter)
    if ((num > int(maxx)) or (num < int(minn))):
        bad += 1
    else:
        good += 1
"""
for x in range(len(full)):

    hy = full[x].find("-")
    co = full[x].find(":")

    minn = int(full[x][0:hy])
    letter = full[x][co-1]
    maxx = int(full[x][(hy+1):(co-2)])
    pw = full[x][(co+2):]

    if (pw[minn-1] == letter) or (pw[maxx-1] == letter):
        if pw[minn-1] == pw[maxx-1]:
            bad += 1
        else:
            good += 1
    else:
        bad += 1


print("bad = " + str(bad))
print("good = " + str(good))