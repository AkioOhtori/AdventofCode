f = open("5.txt")
test = f.read()
f.close()

full = test.splitlines()
seats = []
seatnum = []
missing = []

#First: Establish row - First 7 letters - 0 to 127
#Second: Columnns 0-7 (L/R)
#Third: Seat ID row x 8 + column
#Answer: Highest seat ID

def parse(op, pos):
    if pos[2] <= 1:
        return pos

    pos[2] = int(pos[2]/2)
    if (op == "R") or (op == "B"): #upper half
        pos[0] = pos[2] + pos[0]
    elif (op == "L") or (op == "F"): #lower half
        pos[1] = pos[1] - pos[2]
    else:
        print("!!!!ERROR!!!!")
    return pos

#generate a list of all possible seats
def gen():
    x = 128
    y = 8

    a = []
    while x > 0:
        x = (x - 1)
        while y > 0:
            y = y -1
            a.append((x * 8) + y)
        y = 8
    a.sort(reverse=True)
    return a

seats = gen()

for x in range(len(full)):
    column = [0, 127, 128]
    row = [0, 7, 8]
    for y in range(len(full[x])):
        if y < 7:
            parse(full[x][y], column)
        else:
            parse(full[x][y], row)
    temp = column[0] * 8 + row[0]
    #print("Iteration " + str(x) + " = " + str(temp))
    seatnum.append(temp)


seatnum.sort(reverse = True)
print("\nThe top seat number is " + str(seatnum[0]))

#compile list of missing seats
for x in range(len(seats)):
    if seats[x] not in seatnum:
        missing.append(seats[x])
#compile list of missing seats that have adjacent seats
for x in range(len(missing)):
    if (((missing[x] +1) in seatnum) and ((missing[x] -1) in seatnum)):
        print("Your seat is " +str(missing[x]) + "\n")