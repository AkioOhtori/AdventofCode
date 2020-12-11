f = open("Day11\\11.txt")
test = f.read()
f.close()
temp = test.splitlines()
total_rows = len(temp)
total_columns = len(temp[0])
seatmap = temp.copy()
'''
for y in range(total_rows):
    for x in range(total_columns):
        seats.append(d(temp[y][x])

print(seats)'''

occupied = "#"
floor = "."
empty = "L"

# If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
# If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
# Otherwise, the seat's state does not change.

def is_occupied(r, c, s):
    # y = row
    # x = column
    starty = -1
    startx = -1
    full = 0
    if r == 0: starty = 0
    if c == 0: startx = 0
    for x in range(startx, 2):
        for y in range(starty, 2):
            try:
                if(s[y+r][x+c] == occupied): full += 1
            except:
                continue#print("poop")
    return full

row = 0
column = 0
#for rows, for columns, etc
full_counter = 0
full_counter_old = -1


def tea_party(seats):
    global full_counter
    global full_counter_old

    full_counter_old = full_counter

    seats_old = seats.copy()
    for row in range(len(seats)):
        for column in range(len(seats[row])):
            #this checks for seat adjacency and returns the number of full, adjacent seats
            adjacent = is_occupied(row,column, seats_old)
            if (seats_old[row][column] == empty and (adjacent == 0)):
                    seats[row] = seats[row][:column] + occupied + seats[row][column+1:]
                    full_counter += 1
                    #print("Sat down!")
            elif (seats_old[row][column] == occupied and (adjacent > 4)): #4 is the limit but we always count ourselves
                seats[row] = seats[row][:column] + empty + seats[row][column+1:]
                #print("left!")
                full_counter += -1    
    return seats #return the new list!

while full_counter != full_counter_old:
    seatmap = tea_party(seatmap)
    test = 0
    '''for x in range(len(seats)): test += seats[x].count(occupied)
    if test != full_counter:
        print("PROBLEM!")
        break'''




print(full_counter)
print(seatmap)