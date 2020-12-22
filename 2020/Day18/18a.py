def findparens(l,s):
    return

def domath(l):
    index = 0
    n1 = 0
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
            ret = (domath(l[index:]))
            index += ret[0]
            if op == "": n1 = ret[1]
            else:
                n2 = ret[1]
                break
        elif l[index] == ")":
            index += 1
            return [index, n1]
        else:
            if n1: n2 = int(l[index])
            else: n1 = int(l[index])
            index += 1
            break

    if op == "*":
        a = int(n1) * int(n2)
        print("Did %s * %s and returned %s with index %s" % (n1, n2, a, index))
    else:
        a = int(n1) + int(n2)
        print("Did %s + %s and returned %s with index %s" % (n1, n2, a, index))

    return [index, a]


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
        ret = (domath(line[index:]))
        index += ret[0]
        answer = ret[1]
    total += answer
    print(answer)
print(total)




