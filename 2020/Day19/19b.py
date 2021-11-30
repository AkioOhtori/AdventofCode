import random
import time
start_time = time.time()
f = open("Day19\\19.txt")

rules = dict()

#If a rule contains an or, split it into two new rules and check them
def dealwithor(rl, l, i):
    loc_or = rl.index("|")
    #split the rules into two
    a = rl[:loc_or]
    b = rl[loc_or+1:]
    or1 = "or"+str(random.randrange(1,999999))+"a" #this may not be needed anymore
    or2 = "or"+str(random.randrange(1,999999))+"b"
    rules[or1] = a
    rules[or2] = b
    # HELLO!  REVERSING THE ORDER HERE PRODUCES THE NUMBER OF ERRONIOUS ANSWERS
    # IF YOU RUN OR1 FIRST YOU GET ALMOST THE RIGHT ANSWER, WITH A FEW ODDBALLS
    # IF YOU RUN OR2 FIRST YOU GET ONLY THE ODDBALLS
    # I HAVE LITERALLY NO IDEA WHY THIS IS, BUT SUBTRACT THE TWO AND YOU GET THE REAL ANSWER
    # It seems to be tied to a 8 -> 11 loop?
    inew = checkrule(or1, l, i)
    if inew == 0: #If the first one fails, try the next
        inew = checkrule(or2, l, i)
        if inew == 0:  return 0 #fail

    return inew

#This is a house of cards... possibly literally
def checkrule(r, l, i):
    #If the rule contains an or treat it differently
    if i >= len(l):
        #print("We got to the bad place. i=%s, r=%s and l=%s" % (i,r,len(l)))
        return i
    if rules[r].count("|"):
        inew = dealwithor(rules[r], l, i)
        if inew == 0: return 0
        else: i = inew
    else:
        #Check if we've reached an end state (a or b) or need to go deeper
        for rule in rules[r]:
            if rule == "a":
                if l[i] == "a":
                    i += 1
                else: return 0
            elif rule == "b":
                if l[i] == "b":
                    i += 1
                else: return 0
            else:
                inew = checkrule(rule, l, i)
                if inew == 0: return 0
                else: i = inew
    
    return i #If succeed, return the new index
            

#Preface: Extract Rules and store in a dict
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

rules["8"] = ["42","|","42","8"]
rules["11"] = ["42","31", "|", "42", "11", "31"]


while 1: #main loop!
    line = f.readline().strip()
    if not line: break #make sure there is a empty line at end of input file
    inew = checkrule("0", line, 0)
    if inew != 0 and inew == len(line):
        answer += 1
        #print(line)
print(answer)
print(("%.4f" % (time.time() - start_time)))