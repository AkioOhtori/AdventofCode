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

def thisprobablywontwork(m):
    dump = []
    for x in range(len(m)):
        if m[x] == dontcare:
            dump.append(string_rep(m, "0", x))
            dump.append(string_rep(m, "1", x))
    return dump

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

        for p in range(len(mask)): #Loop through operation, memory, and mask
            if mask[p] == dontcare: memaddr = string_rep(memaddr, dontcare, p)
            elif mask[p] == "1": memaddr = string_rep(memaddr, "1", p)
            #else: memaddr remains unchanged
        memdump = []
        loop = memaddr.count(dontcare)
        #print(loop)
        memdump.append(memaddr)#thisprobablywontwork(memaddr))
        #print(memdump)
        memdump_new = []
        
        while loop > 0:
            memdump_new.clear()
            for mem in memdump:
                memdump_new.extend(thisprobablywontwork(mem))
            memdump.clear()
            memdump = memdump_new.copy()
            loop += -1
        for addr in memdump:
            a = int(addr, 2)
            memory[a] = operation

#print("\n")
print(len(memory))
total = 0


# for addr in memory:
#     n = addr.count(dontcare)
#     total += int(memory[addr])**n
for value in memory: total += int(memory[value]) #Loop through and add up memory

print("\nThe sum of the memory is %s, which took %s seconds to compute.\n" % (total, ((time.time() - start_time))))
