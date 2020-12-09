f = open("9.txt")

preamble = 25 #number of previous instructions per rules
previous = [] #list of previous [preamble] instructions
valid = [] #container for valid next lines, per rules

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

while 1:
    #Generate new table of valid next numbers
    valid = create_valid(previous)
    #Read next line from file
    nextinstruction = int(f.readline().strip())

    #Check for validity
    if nextinstruction not in valid:
        print("The imposter was " + str(nextinstruction))
        break
    previous.append(nextinstruction) #add to rolling previous instructions
    previous.pop(0)  #remove oldest from previous instructions

f.close() #clean-up
