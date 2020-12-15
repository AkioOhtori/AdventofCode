
binstyle = '036b'
default = format(0, binstyle)#"000000000000000000000000000000000000"
memory = dict() #What could possibly go wrong!?
dontcare = "X"
f = "Day14\\14.txt"

mem = ""
mask = ""
operation = ""

def string_rep(s, r, a):
    if r == "": print("ERROR")
    snew = s[:a] + r + s[1+a:]
    if len(snew) != 36: print("ERROR")
    #print("Replacing %s with %s which makes %s" % (s,r,snew))
    return snew

for instruction in open(f):
    instruction = instruction.strip()
    eq = instruction.find("=") +2
    if "mask" in instruction:
        #process old task
        #start new task
        
        mask = instruction[eq:]
        print("New mask is %s" % mask)
    else:
        memend = instruction.find("]")
        membeg = 4
        memaddr = instruction[membeg:memend]
        operation = format(int(instruction[eq:]), binstyle)
        print("New operation is %s going to %s" % (operation, memaddr))

        if memaddr not in memory: memory[memaddr] = default
        
        mem = memory[memaddr]

        for p in range(len(mask)):
            if mask[p] == dontcare: mem = string_rep(mem, operation[p], p)
            else: mem = string_rep(mem, mask[p], p)

        memory[memaddr] = mem
        print("New mem is %s" % mem)

total = 0
for value in memory:
    total += int(memory[value], 2)

print(int(total))
