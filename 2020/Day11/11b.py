import time

fi = open("Day11\\11.txt")
test = fi.read()
fi.close()
seatmap = test.splitlines()

occupied = "#"
floor = "."
empty = "L"

# If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
# If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
# Otherwise, the seat's state does not change.

def increment(n):
    if n == 0:
        return 0
    elif n < 0:
        return (n-1)
    else:
        return (n+1)

def lookaround(x, y, r, c, s, f):
    if ((x+c) < 0 or (y+r) < 0): return f
    try: #if index it out of range deal with it
        if(s[y+r][x+c] == occupied): f += 1
        #elif (s[y+r][x+c] == empty): free += 1
        elif (s[y+r][x+c] == floor):
            xx = increment(x)
            yy = increment(y)
            f = lookaround(xx, yy, r, c, s, f)
    except:
        pass
    return f

def is_occupied(r, c, s): #checks to see how many adjacent seats are occupied
    # y = row
    # x = column
    starty = -1
    startx = -1
    full = 0
    #this is a messy way to deal with negative indexes but meh
    #if r == 0: starty = 0
    #if c == 0: startx = 0
    for x in range(startx, 2):
        for y in range(starty, 2):
            full = lookaround(x, y, r, c, s, full)
    #loop through 8 adjacent cells plus 0,0 - probably a better way to do this
    return full

full_counter = 0
full_counter_old = -1

def tea_party(seats):
    global full_counter
    global full_counter_old

    full_counter_old = full_counter

    old_seating = seats.copy()
    for row in range(len(seats)):
        for column in range(len(seats[row])):
            #this checks for seat adjacency and returns the number of full, adjacent seats
            adjacent = is_occupied(row, column, old_seating)
            if (old_seating[row][column] == empty and (adjacent == 0)):
                #Seat is empty and there are no people around, sit on down!
                seats[row] = seats[row][:column] + occupied + seats[row][column+1:]
                full_counter += 1
            elif (old_seating[row][column] == occupied and (adjacent > 5)): #4 is the limit but we always count ourselves
                #Too many people around!  Time to leave!
                seats[row] = seats[row][:column] + empty + seats[row][column+1:]
                full_counter += -1    
    return seats #return the new list

#Keep changing seats until we settle
start_time = time.time()
while full_counter != full_counter_old: seatmap = tea_party(seatmap)

print("\nThe number of occupied seats is: " + str(full_counter))
print("Execution took %s seconds\n" % (time.time() - start_time))
print(seatmap)
#EOF