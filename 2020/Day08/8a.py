f = open("8.txt")
t = f.read()
f.close()
boot = t.splitlines() #load all instructions into list

#initialize variables because only sometimes python is nice about this
accumulator = 0
instruction = []
pointer = 0
pointer_next = 0
done = 0

#start by parsing "boot" into "instruction"
for x in range(len(boot)):
    i = boot[x][0:3]
    n = boot[x][4:]
    if n.isnumeric: n = int(n)
    else: print("ERROR")
    temp = [i,n]
    instruction.append(temp)

#then execute the code
# acc = increase/decrease accumulator, execute next instruction
# jmp = jump to instruction based on number
# nop = go to next

i = 0 #instruction
n = 1 #number of hops
havedone = 999 #flag this operation has completed before

while 1: #loop though instructions until we repeat
    #print(instruction[pointer])
    if instruction[pointer][i] == havedone: break
    elif instruction[pointer][i] == "acc":
        accumulator += instruction[pointer][n]
        instruction[pointer][i] = havedone
        pointer += 1
    elif instruction[pointer][i] == "jmp":
        instruction[pointer][i] = havedone
        pointer += instruction[pointer][n]
        
    elif instruction[pointer][i] == "nop":
        instruction[pointer][i] = havedone
        pointer += 1
    else: print("ERROR!!")

print("\nAccumulator value was " +str(accumulator) + " when loop detected.\n")