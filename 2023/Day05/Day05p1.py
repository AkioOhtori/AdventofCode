seeds = {}
maps = {}
loc = []
f = open("Day05\\data.txt")

alm = f.read().strip().splitlines()

temp = alm[0][alm[0].find(":")+2:].split(" ")

for seed in temp: seeds[int(seed)] = [int(seed)]

for x in range(2,len(alm)):
    if not alm[x].replace(" ","").isnumeric(): #if we're on heading or whitespace, change dict
        d = alm[x].strip()
        maps[d] = []
        continue
    else:
        m = alm[x].strip().split(" ")
        maps[d].append([int(m[0]), int(m[0])+int(m[2])-1 ,int(m[1]) ,int(m[1])+int(m[2])-1])
maps.pop("") #delete junk dict

for seed in seeds: #iterating over the seeds - first step
    n = 0
    for stage in maps: #iterate over each stage
        for map in maps[stage]: #iterate over each map within a stage [d1, d2, s1, s2]
            if seeds[seed][n] >= map[2] and seeds[seed][n] <= map[3]: #we're on a corresponding recipe
                seeds[seed].append(map[0] + (seeds[seed][n] - map[2]))
                break
        else:
            seeds[seed].append(seeds[seed][n])
        n += 1
    loc.append(seeds[seed][n])
loc.sort()
print(loc[0])