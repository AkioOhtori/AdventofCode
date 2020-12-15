import time
start_time = time.time()
binstyle = '036b'
default = format(0, binstyle)
memory = dict() #What could possibly go wrong!?
dontcare = "X"
f = "Day14\\14.txt"

#Slices string to insert new character (you can't do this like a list because python!?)
def string_rep(s, r, a):
    snew = s[:a] + r + s[1+a:]
    return snew

for instruction in open(f):
    instruction = instruction.strip()
    eq = instruction.find("=") +2
    if "mask" in instruction:
        mask = instruction[eq:]
        #print("New mask is %s" % mask)
    else:
        #Locate reference points for string slicing
        memend = instruction.find("]")
        membeg = 4 #hard coded as it doesn't change
        memaddr = instruction[membeg:memend] #isolate memory address
        operation = format(int(instruction[eq:]), binstyle) #Isolation opearation and format to 36bit binary
        #print("New operation is %s going to %s" % (operation, memaddr))

        if memaddr not in memory: memory[memaddr] = default #create dir entry        
        mem = memory[memaddr] #store current memory in variable

        for p in range(len(mask)): #Loop through operation, memory, and mask
            if mask[p] == dontcare: mem = string_rep(mem, operation[p], p)
            else: mem = string_rep(mem, mask[p], p)

        memory[memaddr] = mem #store new value of memory
        #print("New mem is %s" % mem)

total = 0
for value in memory: total += int(memory[value], 2) #Loop through and add up memory

print("\nThe sum of the memory is %s, which took %s seconds to compute.\n" % (total, ((time.time() - start_time))))
