f = open("8.txt")
t = f.read()
f.close()
boot = t.splitlines()

acc = 0
inst = []
point = 0



#start by parsing "boot"
for x in range(len(boot)):
    i = boot[x][0:3]
    n = boot[x][4:]
    if n.isnumeric: n = int(n)
    else: print("ERROR")
    temp = [i,n]
    inst.append(temp)

#then execute the code
# acc = increase/decrease accumulator, execute next instruction
# jmp = jump to instruction based on number
# nop = go to next

def runout(ins, pointer, accumulator):
    done = 0
    i = 0 #instruction
    n = 1 #number of hops
    havedone = 999
    while not done:
        print(ins[pointer])
        #if pointer >= len(ins):
        if ins[pointer][i] == havedone:
            done = 1
            break
        elif ins[pointer][i] == "acc":
            accumulator += ins[pointer][n]
            ins[pointer][i] = havedone
            pointer += 1
        elif ins[pointer][i] == "jmp":
            ins[pointer][i] = havedone
            pointer += ins[pointer][n]
            
        elif ins[pointer][i] == "nop":
            ins[pointer][i] = havedone
            pointer += 1
        else:
            print("ERROR")
    print(accumulator)

runout(inst, point, acc)