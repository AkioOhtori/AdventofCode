import random
f = open("Day19\\19.txt")

rules = dict()

def dealwithor(rl, l, i):
    loc_or = rl.index("|")
    a = rl[:loc_or]
    b = rl[loc_or+1:]
    or1 = "or"+str(random.randrange(1,999999))+"a"
    or2 = "or"+str(random.randrange(1,999999))+"b"
    rules[or1] = a
    rules[or2] = b
    inew = checkrule(or1, l, i)
    if inew == 0:
        inew = checkrule(or2, l, i)
        if inew == 0:
            return 0
        else: return inew
    else:
        return inew



def checkrule(r, l, i):
    #print("Checking rule %s with index %s" % (r,i))
    rule_list = rules[r]
    if rule_list.count("|"):
        inew = dealwithor(rule_list, l, i)
        if inew == 0: return 0
        else: i = inew
    else:
        for rule in rule_list:
            if rule == "a":
                if l[i] == "a":
                    i += 1
                else:
                    #print("Failed at index %s on rule %s" % (i, r))
                    return 0
            elif rule == "b":
                if l[i] == "b":
                    i += 1
                else:
                    #print("Failed at index %s on rule %s" % (i, r))
                    return 0
            else:
                inew = checkrule(rule, l, i)
                if inew == 0:
                    #print("Failed at index %s on rule %s" % (i, r))
                    return 0
                else: i = inew
    return i
            

    

#Extract Rules
while 1:
    line = f.readline().strip()
    if line != "":
        rule_numer = ""
        colon = line.index(":")
        rule_numer = line[:colon]
        rule = line[colon+2:].strip("\"").split(" ")
        rules[rule_numer] = rule
    else: break

answer = 0

while 1: #main loop!
    line = f.readline().strip()
    if not line: break #make sure there is a empty line at end of input file
    # print(line)
    inew = checkrule("0", line, 0)
    if inew != 0 and inew == len(line):
        #print("%s GOOD. Index %s" % (line, inew))
        answer += 1
    #break
    # else: print(line)
print(answer)