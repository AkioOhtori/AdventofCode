import numpy as np
import time
start_time = time.time()
fname = "Day20\\20.txt"
length = 10
width = 10
tiles = dict()
tile_sides = dict()
top = 0
bottom = 1
left = 2
right = 3
a = []

#Start by converting all the tiles to matricies and storing them in a dict
for line in open(fname):
    line = line.strip()
    if not line:
        #Done with this tile, time to store it
        temp = np.array(a, copy=True)
        temp = temp.reshape(length,width).copy()
        tiles[title]=temp.copy()
        tile_sides[title] = [0,0,0,0]
        title = ""
        a.clear()
    elif "Tile" in line:
        c = line.find(":")
        title = line[c-4:c]
    else:
        for p in line:
            if p == "#": a.append(1)
            else: a.append(0)

#don't actually need to reconstruct the image... just need to find 4 tiles that don't have matches on two sides

def checkfour(m1, m2, tile, side):
    if tile == other: return
    if np.array_equal(m1[0], m2[0]): #Top to top
        tile_sides[tile][side] +=1
    if np.array_equal(m1[0], m2[-1]): #top to bottom
        tile_sides[tile][side] +=1
    if np.array_equal(m1[-1], m2[0]): #Bottom to top
        tile_sides[tile][side+1] +=1
    if np.array_equal(m1[-1], m2[-1]): #Bottom to bottom
        tile_sides[tile][side+1] +=1

#Yes, there are better ways to do this but heh
for tile in tiles:
    for other in tiles:
        #Check the primary tile against the others in all configurations
        checkfour(tiles[tile], tiles[other], tile, top)
        checkfour(tiles[tile], np.fliplr(tiles[other]), tile, top)
        checkfour(tiles[tile], np.swapaxes(tiles[other],1,0), tile, top)
        checkfour(tiles[tile], np.fliplr(np.swapaxes(tiles[other],1,0)), tile, top)
        checkfour(np.swapaxes(tiles[tile],1,0), tiles[other], tile, left)
        checkfour(np.swapaxes(tiles[tile],1,0), np.fliplr(tiles[other]), tile,left)
        checkfour(np.swapaxes(tiles[tile],1,0), np.swapaxes(tiles[other],1,0), tile, left)
        checkfour(np.swapaxes(tiles[tile],1,0), np.fliplr(np.swapaxes(tiles[other],1,0)), tile, left)

answer = 1

#Check for tiles that only matched on two sides and you have your answer
for tile in tiles:
    if (np.count_nonzero(tile_sides[tile])) == 2:
        answer *= int(tile)

print("The answer is %s and it took %s seconds to compute!" % (answer, ("%.2f" % (time.time() - start_time))))