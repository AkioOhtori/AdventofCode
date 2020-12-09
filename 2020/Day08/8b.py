f = open("8.txt")
t = f.read()
f.close()
boot = t.splitlines()
#Above loads all instructions into list "boot"

#Initialize variables becaues you need to do that sometimes
instruction = []
point = 0
inst_old = str()
ans = []

#start by parsing "boot" into "instruction"
for x in range(len(boot)):
    i = boot[x][0:3]
    n = boot[x][4:]
    if n.isnumeric: n = int(n)
    else: print("ERROR")
    temp = [i,n]
    instruction.append(temp)

#References for the instruction array (has to live here!)
i = 0 #instruction
n = 1 #number of hops

# Rules:
# acc = increase/decrease accumulator, execute next instruction
# jmp = jump to instruction based on number
# nop = go to next

#Runs through the complete instruction set with the replaced line
def runout(ins):
    pointer = 0
    accumulator = 0
    havedone = []
    while 1:
        if pointer in havedone:#len(havedone) > (5*len(ins)): #Tried this two ways, both worked
            return [0, accumulator]
        elif pointer >= len(ins): #HAVE WE ESCAPED!?
            return [1, accumulator]
        elif ins[pointer][i] == "acc":
            accumulator += ins[pointer][n]
            havedone.append(pointer)
            pointer += 1
        elif ins[pointer][i] == "jmp":
            havedone.append(pointer)
            pointer += ins[pointer][n]
        elif ins[pointer][i] == "nop":
            havedone.append(pointer)
            pointer += 1
        else: print("ERROR")

def adv_pointer(line, p): #increment p(ointer) as appropriate and return
    if line[i] == "jmp": p += line[n]
    else: p += 1
    return p

# 1st save instruction
# Then sub instruction with nop/jmp
# If program didn't terminate, restore and bump to the next

while 1:  #Loop until a sub that causes an end is found
    inst_old = instruction[point][i]

    if inst_old == "nop":
        instruction[point][i] = "jmp"
    elif inst_old == "jmp":
        instruction[point][i] = "nop"
    #blegh this runs for acc operations but skipping would be messy
    ans = runout(instruction)
    if ans[0]: #This indicates an end was found.  Win!
        print("\nReplaced line " + str(point) + " with " + str(instruction[point][i]))
        print("Accumulator was " + str(ans[1]) + " when terminated\n")
        break

    instruction[point][i] = inst_old #since no answer was found, restore that instruction
    point = int(adv_pointer(instruction[point], point)) #update pointer to next

exit #not required but it felt weird to not have

