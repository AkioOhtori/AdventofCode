# Adapter rating = +0j, - 1-3j
# Build in adatper = highest + 3j
# Outlet = 0j
name = "Day10\\10as1.txt"
f = open(name)
test = f.read().splitlines()
f.close()
adapters = [0] #0 for outlet
arragements = 0
iterations = 0

for a in range(len(test)): adapters.append(int(test[a].strip()))

adapters.append(adapters[len(adapters)-1] + 3) #append internal adapter
adapters.sort(reverse=True) #sort smallest to largest
end = len(adapters)-1

ways_to = dict()

for x in adapters: ways_to[x] = 0

for x in adapters:
    print(x)
    y = 0
    for n in range(0,4):
        if (x + n) <= end:
            if (adapters[x] - adapters[x+n]) <= 3:
                y += 1
    ways_to[x] += y

answer = 0
for x in adapters:
    answer = answer+ways_to[x]
print(ways_to)
print(answer)

#EOF
