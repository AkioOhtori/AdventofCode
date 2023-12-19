import time
start_time = time.time()

seeds = {}
maps = {}

f = open("Day05\\taylor.txt")
alm = f.read().strip().splitlines()

temp = alm[0][alm[0].find(":")+2:].split(" ") #This creates a clean seed list
for x in range(0, len(temp), 2):
    seeds[int(temp[x])] = [int(temp[x]),int(temp[x])+int(temp[x+1])-1]
    # for part 2 we're formatting seeds into a seed dict with seed: [min_val, max_value]
    # the index doesn't actually matter and this could be a list instead but heh

# This is the same as part 1 - take the rest of the text and shove it in a dict
# no it doesn't NEED to be a dict but it makes things auto-generate and that makes me happy
for x in range(2,len(alm)):
    if not alm[x].replace(" ","").isnumeric(): #if we're on heading or whitespace, change entry
        d = alm[x].strip()
        maps[d] = []
        continue
    else: #otherwise we're converting the X Y Range of the input data to [Xmin, Xmax, Ymin, Ymax]
        m = alm[x].strip().split(" ")
        maps[d].append([int(m[0]), int(m[0])+int(m[2])-1 ,int(m[1]) ,int(m[1])+int(m[2])-1])
maps.pop("") #delete white-space junk entry
"""
    stage/dict name       map1              map2
maps {                    min max min max
    'seed-to-soil map:': [[50, 51, 98, 99], [52, 99, 50, 97]], 
    'soil-to-fertilizer map:': [[0, 36, 15, 51], [37, 38, 52, 53], [39, 53, 0, 14]], 
    'fertilizer-to-water map:': [[49, 56, 53, 60], [0, 41, 11, 52], [42, 48, 0, 6], [57, 60, 7, 10]], 
    'water-to-light map:': [[88, 94, 18, 24], [18, 87, 25, 94]], 
    'light-to-temperature map:': [[45, 67, 77, 99], [81, 99, 45, 63], [68, 80, 64, 76]], 
    'temperature-to-humidity map:': [[0, 0, 69, 69], [1, 69, 0, 68]], 
    'humidity-to-location map:': [[60, 96, 56, 92], [56, 59, 93, 96]]}
As mentioned above, we're expanding out the ranges for easier compare, though it could be done either way
"""

#reverse dictionary because we're coming in from behind (heyo) for part 2
res = dict(reversed(list(maps.items())))

loc = 1 #starting location; Could comfortably be in the millions but why not here
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
            print("\nThe lowest starting location is %s from seed %s, which took %s seconds to compute.\n" % (loc, next, ("%.5f" % (time.time() - start_time))))
            loc = 50000000001 #force ourselves out of the loop
            break