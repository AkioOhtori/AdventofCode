
#Takes in any list containing a + and resolves the first one
def doPluslist(l):
    lnew = []
    s = l.index("+")
    a = int(l[s-1]) + int(l[s+1])
    lnew = l[:s-1] + [a] + l[s+2:]
    return lnew

#Takes in any list containing a * and resolves the first one
def doMultlist(l):
    lnew = []
    s = l.index("*")
    a = int(l[s-1]) * int(l[s+1])
    lnew = l[:s-1] + [a] + l[s+2:]
    return lnew

#Takes in any list containing an expression in parens
#resolves the innermost parens and returns a new list
def doParenslist(l):
    lnew = []
    e1 = l.index(")")
    s1 = e1
    while s1 >= 0:
        if l[s1] == "(":
            break
        else: s1 += -1
    lnew = l[s1+1:e1]
    # Completely resolve all math in the parens
    while lnew.count("+"): lnew = doPluslist(lnew)
    while lnew.count("*"): lnew = doMultlist(lnew)
    
    lnew =l[:s1] + lnew + l[e1+1:]
    #return the list with the a paren resolved to a number
    return lnew

total = 0
for line in open("Day18\\18.txt"):
    line = line.strip()
    line = line.replace(" ","")
    llist = []
    for n in line: llist.append(n)

    #resolve all parens, then addition, then multiplication
    while llist.count("("): llist = doParenslist(llist)
    while llist.count("+"): llist = doPluslist(llist)
    while llist.count("*"): llist = doMultlist(llist)
    total += int(llist[0])
print("\nThe sum of all the math problems is %s\n" % total)
#EOF