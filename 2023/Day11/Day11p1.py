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

#expand
index = 0
for y in rows:
    image.insert(y+index, "."*len(image[0]))
    index += 1

index = 0
for x in cols:
    for y in range(len(image)):
        image[y] = image[y][:x+index] + "." + image[y][x+index:]
    index += 1

#enumerate galaxies
    
for y in range(len(image)):
    for x in range(len(image[y])):
        if image[y][x] == "#": gal.append([y,x])

for c in gal:
    for c2 in gal:
        if c != c2:
            a = abs(c[0]-c2[0])
            b = abs(c[1]-c2[1])
            d += a+b

print(int(d/2))