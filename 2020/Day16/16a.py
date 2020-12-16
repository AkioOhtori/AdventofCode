fields = dict()
valid = set()
myticket = []
validctr = 0
invalidcrt = 0
# References for range lists
ll = 0 #Low range low limit
lh = 1 #Low range high limit
hl = 2 #High range low limit
hh = 3 #High range high limit

#Alright lets take a look at this
f = open("Day16\\16.txt")

#LOL jk functions first!

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
    updatevalid(rangelist) #Go ahead and add this lists range to valid set
    return temp

#Split CSV string into list of ints
def parsecsv(line):
    output = []
    temp = (line.split(","))
    for x in range(len(temp)): output.append(int(temp[x]))
    return output

#Feed me a range and I'll update valaid list
def updatevalid(r):
    global valid
    for x in range(r[ll], r[lh]+1): valid.add(x) #lower range
    for x in range(r[hl], r[hh]+1): valid.add(x) #upper range
    return 0

#Read initial data
while 1:
    l = f.readline().strip()
    if not l: continue
    elif "your ticket" in l:
        l = f.readline().strip()
        myticket = parsecsv(l)
    elif "nearby" in l: break
    else: fields.update(parsefield(l))

#Should in initialized at this point.  Now what!?
#The next line we read will be the next list of fields on a ticket
#So read each line and add up the invalid items

while 1:
    l = f.readline().strip()
    if not l: break #if we've reached the end, end
    line = parsecsv(l)
    for n in line:
        if n not in valid: invalidcrt += n

print("\n The sum of the invalid fields was %s!\n" % invalidcrt)
f.close()
#EOF
