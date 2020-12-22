def doPluslist(l):
    lnew = []
    s = l.index("+")
    a = int(l[s-1]) + int(l[s+1])
    lnew = l[:s-1] + [a] + l[s+2:]
    return lnew

def doMultlist(l):
    lnew = []
    s = l.index("*")
    a = int(l[s-1]) * int(l[s+1])
    lnew = l[:s-1] + [a] + l[s+2:]
    return lnew


def doParenslist(l):
    lnew = []
    e1 = l.index(")")
    s1 = e1
    while s1 >= 0:
        if l[s1] == "(":
            break
        else: s1 += -1
    lnew = l[s1+1:e1]
    # print(lnew)
    while lnew.count("+"): lnew = doPluslist(lnew)
    while lnew.count("*"): lnew = doMultlist(lnew)
    
    lnew =l[:s1] + lnew + l[e1+1:]
    # print(lnew)
    return lnew#[lnew, s1, e1]

def doPlus(l):
    
    s1 = l.find("+")
    n1 = ""#l[s1-i]
    n2 = ""#l[s1+i]
    i1 = 1
    try:
        while l[s1-i1].isnumeric and (s1-i1) >= 0:
            n1 = l[(s1-i1)] + n1
            i1 += 1
    except:
        pass
    print("n1 is %s" % n1)
    i2 = 1
    test = l[s1+i2].copy()
    while (s1+i2) < len(l):
        if test.isdecimal:
            n2 = n2 + l[s1+i2]
            i2 += 1
            test = l[s1+i2].copy()
        else:
            break
    print("n2 is %s" % n2)
    a = 0
    a = int(n1.strip()) + int(n2.strip())
    
    lnew = l[:s1-i1] + str(a) + l[s1+i2:]
    return lnew

total = 0
for line in open("Day18\\18.txt"):
    line = line.strip()
    line = line.replace(" ","")
    llist = []
    for n in line: llist.append(n)

    while llist.count("("): llist = doParenslist(llist)
    while llist.count("+"): llist = doPluslist(llist)
    while llist.count("*"): llist = doMultlist(llist)
    total += int(llist[0])
    print(llist[0])
print(total)
