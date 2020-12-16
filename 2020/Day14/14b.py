import time
start_time = time.time()
binstyle = '036b'
default = format(0, binstyle)
memory = dict() #What could possibly go wrong!?
dontcare = "X"
f = "Day14\\14.txt"

for instruction in open(f):
    instruction = instruction.strip()
    eq = instruction.find("=") +2
    if "mask" in instruction:
        mask = instruction[eq:]
    else:
        #Locate reference points for string slicing
        memend = instruction.find("]")
        membeg = 4 #hard coded as it doesn't change
        memaddr = format(int(instruction[membeg:memend]), binstyle) #isolate and format memory address
        operation = int(instruction[eq:])

        newmem = [""] #container for memory updates
        for p in range(len(mask)): #Loop through operation, memory, and mask
            if mask[p] == dontcare: #need to add new to the string PLUS make a new set
                z = len(newmem) #since length will change, need to store existing length
                for x in range(0,z): #Iterate over OG list
                    newmem.append(newmem[x]+"1") #Add the new line first
                    newmem[x] += "0" #then continue with old
            elif mask[p] == "1":
                for x in range(len(newmem)): newmem[x] += "1" #cat 1 to every member of the list
            else: 
                for x in range(len(newmem)): newmem[x] += memaddr[p] #cat addr to every

        for x in newmem: memory[int(x,2)] = operation #update memory

total = 0
for value in memory: total += int(memory[value]) #Loop through and add up memory

print("\nThe sum of %s memory items is %s, which took %s seconds to compute.\n" % (len(memory), total, ("%.5f" % (time.time() - start_time))))
