# Adapter rating = +0j, - 1-3j
# Build in adatper = highest + 3j
# Outlet = 0j
name = "Day10\\10.txt"
f = open(name)
test = f.read().splitlines()
f.close()
adapters = list()#[0] #0 for outlet
adapters.append(0)
arragements = 0
iterations = 0

for a in range(len(test)): adapters.append(int(test[a].strip()))

adapters.sort(reverse=True) #sort smallest to largest
adapters.insert(0, adapters[0] + 3) #append internal adapter
end = len(adapters)-1

ways_to = dict()
#print(adapters)

for x in adapters: ways_to[x] = 0

ways_to[adapters[0]] = 1

for x in range(len(adapters)):
    y = ways_to[adapters[x]]
    print("Adapter " + str(adapters[x]) + ", dict: " + str(ways_to[adapters[x]]))
    for n in range(1,4):
        if (x + n) <= end:
            if (adapters[x] - adapters[x+n]) <= 3:
                ways_to[adapters[x+n]] += y

#answer = 0
#for x in adapters:
#    answer = answer+ways_to[x]
print("\nThere are " + str(ways_to[0]) + " combinations that will work.\n")
#print(answer)

#EOF
