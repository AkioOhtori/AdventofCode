import numpy as np

map = []
steps = 0
loc = []
dist = 0
ans = 0

GO = {
    "|": "NS",
    "-": "EW",
    "L": "NE",
    "J": "NW",
    "7": "WS",
    "F": "ES",
    ".": "",
    "S": "NSEW"
}
VALID = { #from direction into shape
    "N": "|S7F",
    "S": "|SJL",
    "W": "S-LF",
    "E": "S-J7"
}
OPPO = {
    "N": "S",
    "S": "N",
    "E": "W",
    "W": "E"
}
dir = {
    "N": [-1,0],
    "S": [1,0],
    "E": [0,1],
    "W": [0,-1]
}

"""
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
"""

# def direction(d, n):
#     #
#     x = GO[n] - 
#     return new

for line in open("Day10\\data.txt"):
    map.append(line.strip())
    if "S" in line:
        s = [len(map) - 1, line.find("S")]

for direction in GO["S"]:
    loc = s
    dist = 1
    loc = np.array(loc) + np.array(dir[direction])
    next = map[loc[0]][loc[1]]
    while next in VALID[direction] and next != "S":
        direction = GO[next].replace(OPPO[direction],"") #the direction coming out of the transformation
        #now we need to calculate the new loc
        loc = np.array(loc) + np.array(dir[direction])
        next = map[loc[0]][loc[1]]
        dist += 1
        if next == "S":
            ans = dist/2
            print(ans)
            break


# while loc != "S" or loc != ".":
