fields = dict()
valid = set()
myticket = []
validctr = 0
invalidcrt = 0
ll = 0
lh = 1
hl = 2
hh = 3
f = open("Day16\\16.txt")

#find and parse fields - MESSY! T_T
def parsefield(line):
    colon = line.find(":")
    ranges = line[colon+2:]
    fieldname = line[:colon]

    mid = ranges.find("or")
    lower = ranges[:mid-1]
    upper = ranges[mid+3:]

    hyphen = lower.find("-")
    rangelist = []
    rangelist.append(int(lower[:hyphen]))
    rangelist.append(int(lower[hyphen+1:]))
    hyphen = upper.find("-")
    rangelist.append(int(upper[:hyphen]))
    rangelist.append(int(upper[hyphen+1:]))

    temp = dict()
    temp[fieldname] = rangelist
    updatevalid(rangelist)
    return temp

def parsecsv(line):
    output = []
    temp = (line.split(","))
    for x in range(len(temp)): output.append(int(temp[x]))
    return output

#Feed me a range and I'll update invlaid list
def updatevalid(r):
    global valid
    for x in range(r[ll], r[lh]+1): valid.add(x) #lower range
    for x in range(r[hl], r[hh]+1): valid.add(x) #upper range
    return 0


while 1:
    l = f.readline().strip()
    if not l: continue
    elif "your ticket" in l:
        l = f.readline().strip()
        myticket = parsecsv(l)
    elif "nearby" in l: break
    else: fields.update(parsefield(l))
# print(fields)
# print(myticket)
# print(valid)

#Should in initialized at this point.  Now what!?
#The next line we read will be the next list o tickets
#So we turn that into a list, and determine if the item is valid
#That is going to suck...
while 1:
    l = f.readline().strip()
    if not l: break
    line = parsecsv(l)
#    for n in range(len(line)):
    for n in line:
        if n in valid: validctr += n#line[n]
        else: invalidcrt += n#line[n]

print(validctr)
print(invalidcrt)
f.close()