def findparens(l,s):
    return

def domath(l):
    index = 0
    done = 0
    a = 0
    # global index
    
    while not done:
        n1 = a
        op = ""
        n2 = 0
        
        while not n2:
            if l[index] == "*" or l[index] == "+":
                op = l[index]
                index += 1
            elif l[index] == "(":
                index += 1
                ret = []
                ret = (domath(l[index:]))
                index += ret[0]
                if n1 == 0: n1 = ret[1]
                else:
                    n2 = ret[1]
                    break
            elif l[index] == ")":
                index += 1
                done = 1
                return [index, n1]
            else:
                if n1:
                    n2 = int(l[index])
                    index += 1
                    break
                else:
                    n1 = int(l[index])
                    index += 1
        if op == "*":
            a = int(n1) * int(n2)
            # print("Did %s * %s and returned %s with index %s" % (n1, n2, a, index))
        else:
            a = int(n1) + int(n2)
            # print("Did %s + %s and returned %s with index %s" % (n1, n2, a, index))
        if index >= len(l):
            return(index,a)

    # return [index, a]


total = 0
for line in open("Day18\\18.txt"):
    line = line.strip()
    line = line.replace(" ","")

    index = 0
    answer = 0
    while index <= len(line)-1:
        ret = []
        ret = (domath(line[index:]))
        index += ret[0]
        answer = ret[1]
    total += answer
    print(answer)
print(total)




