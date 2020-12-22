def findparens(l,s):
    return

def domath(n1, l):
    index = 0
    s2 = ""
    op = ""
    n2 = 0
    a = 0
    while 1:
        if l[index] == "*" or l[index] == "+":
            op = l[index]
            index += 1
        elif l[index] == "(":
            index += 1
            ret = []
            ret = (domath(0,l[index:]))
            index += ret[0]
            if op == "": n1 = ret[1]
            else:
                n2 = ret[1]
                break
        elif l[index] == ")":
            index += 1
            return [index, n1]
        else:
            s2 = l[index]
            n2 = int(s2)
            index += 1
            break

    if op == "*":
        a = int(n1) * int(n2)
        print("Did %s * %s and returned %s with index %s" % (n1, n2, a, index))
    else:
        a = int(n1) + int(n2)
        print("Did %s + %s and returned %s with index %s" % (n1, n2, a, index))

    return [index+1, a]


total = 0
for line in open("Day18\\18s.txt"):
    line = line.strip()
    line = line.replace(" ","")
    #print(line)
    # f = (line.count(")"))
    s = (line.find("("))+1
    f = (line.find(")"))
    p = line[s:f]
    #are there parens?
    #find them, catch them, find more!
    index = 0
    answer = 0
    # if line[index+1] == "(": domath(line[index+2:])
    while index <= len(line)-1:
        ret = []
        ret = (domath(0,line[index:]))
        index += ret[0]
        answer = ret[1]
    total += answer
    print(answer)
print(total)




