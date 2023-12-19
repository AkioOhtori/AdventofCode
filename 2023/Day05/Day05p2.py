import time
start_time = time.time()

seeds = {}
maps = {}
f = open("Day05\\data.txt")
s = "seed-to-soil map:"
alm = f.read().strip().splitlines()

temp = alm[0][alm[0].find(":")+2:].split(" ")


for x in range(0, len(temp), 2):
    seeds[int(temp[x])] = [int(temp[x]),int(temp[x])+int(temp[x+1])-1]

# This is the same as part 1
for x in range(2,len(alm)):
    if not alm[x].replace(" ","").isnumeric(): #if we're on heading or whitespace, change dict
        d = alm[x].strip()
        maps[d] = []
        continue
    else:
        m = alm[x].strip().split(" ")
        maps[d].append([int(m[0]), int(m[0])+int(m[2])-1 ,int(m[1]) ,int(m[1])+int(m[2])-1])
maps.pop("") #delete junk dict

#reverse dictionary because we're coming in from behind (heyo)
res = dict(reversed(list(maps.items())))

loc = 1 #this will iterate but we need to keep track of it
while loc < 50000000000:
    loc += 1
    next = loc
    for stage in res: #back calculate the seed from the location
        for map in res[stage]:
            if next >= map[0] and next <= map[1]:
                next = map[2] + next - map[0]
                break
    #validate "next" (now as seed) against starting seeds
    for s in seeds:
        if next >= seeds[s][0] and next <= seeds[s][1]:
            print("Winner! ", next, loc)
            print("\nThe lowest starting location is %s from seed %s, which took %s seconds to compute.\n" % (loc, next, ((time.time() - start_time))))
            loc = 50000000001
            break