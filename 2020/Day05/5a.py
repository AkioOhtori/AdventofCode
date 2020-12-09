f = open("5.txt")
test = f.read()
f.close()

full = test.splitlines()

#First: Establish row - First 7 letters - 0 to 127
#Second: Columnns 0-7 (L/R)
#Third: Seat ID row x 8 + column
#Answer: Highest seat ID

def parse(op, pos):
    if pos[2] <= 1: #if we've reached the bottom
        return pos
    pos[2] = int(pos[2]/2) #update the available set
    #slice based on number
    if (op == "R") or (op == "B"): #upper half
        pos[0] = pos[2] + pos[0]
    elif (op == "L") or (op == "F"): #lower half
        pos[1] = pos[1] - pos[2]
    else: print("!!!!ERROR!!!!")
    return pos #Send seat position back

seatnum = []

#I have no idea how or why this works
for x in range(len(full)):
    column = [0, 127, 128]
    row = [0, 7, 8]
    for y in range(len(full[x])):
        if y < 7:
            parse(full[x][y], column)
        else:
            parse(full[x][y], row)
    temp = column[0] * 8 + row[0]
    print("Iteration " + str(x) + " = " + str(temp))
    seatnum.append(temp)

seatnum.sort(reverse = True)
print("The top seat number = " + str(seatnum[0]))
