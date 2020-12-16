import time
start_time = time.time()

fields = dict()
valid = set()
myticket = []
validctr = 0
invalidcrt = 0
woohoo = []
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
    temp[fieldname] = updatevalid(rangelist) #Go ahead and add this lists range to valid set
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
    v = []
    for x in range(r[ll], r[lh]+1): v.append(x) #lower range
    for x in range(r[hl], r[hh]+1): v.append(x) #upper range
    valid.update(v)
    #v.append(0)
    return v

#Read initial data
while 1:
    l = f.readline().strip()
    if not l: continue
    elif "your ticket" in l:
        l = f.readline().strip()
        myticket = parsecsv(l)
    elif "nearby" in l: break
    else: fields.update(parsefield(l))

#Because dicts aren't ordered!?, we need to make an order
fieldnames = list()
for name in fields: fieldnames.append(name)

#Need a fields x positions matrix of 1s
test = [1] * len(myticket)
solutionmatrix = [] #really need a better name for this, but it was "test2" so...
for x in range(len(fields)):
    solutionmatrix.append(test.copy()) #ARG FUCKING COPY PYTHON GARBAGE

#Should in initialized at this point.  Now what!?
#The next line we read will be the next list of fields on a ticket
while 1:
    l = f.readline().strip()
    if not l: break
    ticket = parsecsv(l)
    for position in range(len(ticket)): #go through each number on a ticket
        if ticket[position] not in valid: #if that number isn't valid... rm
            invalidcrt += ticket[position] #part 1 answer things
            break #stop evaluating this ticket
        else: #otherwise number is valid for SOME field
            for name in range(len(fieldnames)): #is number if valid for a given field cool, if not make it 0
                if ticket[position] not in fields[fieldnames[name]]: solutionmatrix[position][name] *= 0
                #else: solutionmatrix[position][x] *= 1 #which is trivial

#OK now we have a matrix of valid answers. 
#This is where things kinda go to shit
#solutionmatrix[position][field]!
answers = dict()
unsolved = 999
while unsolved > 0:
    unsolved = 0 #HEY WE'RE DONE! Oh wait... just initializing for loop
    for position in range(len(solutionmatrix)):
        ok = solutionmatrix[position].count(1)
        if ok == 1: #If this position is solved, fine, overwrite, and store answer
            #yes, this is dumb and messy and I don't care for it ONE BIT
            field = solutionmatrix[position].index(1)
            for y in range(len(solutionmatrix)): solutionmatrix[y][field] = 0
            answers[fieldnames[field]] = position #This will make no sense tomorrow
        elif ok == 0: pass
        else: unsolved += 1

a = 1 #initialize answer
for x in answers:
    if "departure" in x: #find departure fields, as per problem statement
        a *= myticket[answers[x]] #multiply them together, as per problem statement

print("\nThe answer to part one, the sum of the invalid fields, was %s!" % invalidcrt)
print("The answer for part two is %s, which took %s to compute.\n" % (a, ("%.5f" % (time.time() - start_time))))
f.close()
#EOF
