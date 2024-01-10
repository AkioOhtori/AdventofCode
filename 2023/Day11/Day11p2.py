import numpy as np
#enum rows
#enum columns
#expand rows
#expand columns

#enum galaxies
#math

rows = []
cols= []
d = 0
gal = []
done = []
e = 999999
#I think I'll keep them as strings... for now
image = []
for line in open("Day11\\data.txt"):
    image.append(line.strip())

for x in range(len(image[0])): #find empty rows and columns
    col = ""
    if image[x] == ("."*len(image[x])):
        rows.append(x)
    for y in range(len(image)):
        col += image[y][x]
    if col == ("."*len(image)):
        cols.append(x)


#enumerate galaxies
    
for y in range(len(image)):
    for x in range(len(image[y])):
        if image[y][x] == "#": gal.append([y,x])

for c in gal:
    for c2 in gal:
        if c != c2:
            a = abs(c[0]-c2[0]) #y
            for r in rows:
                if c[0] >= r and c2[0] <= r: a += e
                elif c2[0] >= r and c[0] <= r: a += e
            b = abs(c[1]-c2[1]) #x
            for x in cols:
                if c[1] >= x and c2[1] <= x: b += e
                elif c2[1] >= x and c[1] <= x: b += e
            d += a+b

print(int(d/2))