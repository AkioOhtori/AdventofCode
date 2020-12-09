f = open("9.txt")

preamble = 25 #number of previous instructions per rules
previous = [] #list of previous [preamble] instructions
valid = [] #container for valid next lines, per rules
oldnumbers = [] #container for previous executed lines
imposter = 0 #storage for "bad" instruction (not needed)

#fill in the preable
n = 0
while n < preamble:
    temp = f.readline().strip()
    previous.append(int(temp))
    n += 1

#Fill a list with all the valid next numbers
def create_valid(nums):
    ans = []
    for x in range(len(nums)):
        for y in range(len(nums)):
            if nums[x] != nums[y]: #numbers can't be the same, per rules
                ans.append(nums[x] + nums[y])
    return(ans)
while 1: #Check for "bad" instruction, exit when found
    #Generate new table of valid next numbers
    valid = create_valid(previous)
    #Read next line from file
    nextinstruction = int(f.readline().strip())

    #Check for validity
    if nextinstruction not in valid:
        imposter = nextinstruction
        #store previous [preable] commands in "old commands"
        [oldnumbers.append(i) for i in previous]
        print("\nThe imposter was " + str(imposter))
        break
    previous.append(nextinstruction) #load in next instruction
    oldnumbers.append(previous.pop(0)) #remove and store oldest
f.close() #clean up

for z in range(len(oldnumbers)):
    n = 1
    #check for out of range index
    if ((z+n) >= len(oldnumbers)): break
    #initialize addition
    sums = oldnumbers[z] + oldnumbers[z+n]
    while sums < imposter: #keep adding until something happens
        n += 1
        if ((z+n) >= (len(oldnumbers) + 1)): break
        sums += oldnumbers[z+n]
    if sums == imposter: break #WINNER!

answers = [] #initialize
#pull out the numbers that add to answer
for z in range(z, (z+n+1)): answers.append(int(oldnumbers[z]))

answers.sort() #organize numerically
print("The answer is " + str((answers[0] + answers[(len(answers)-1)]))+"\n")
        


