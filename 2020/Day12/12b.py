
'''
Action N means to move north by the given value.
Action S means to move south by the given value.
Action E means to move east by the given value.
Action W means to move west by the given value.
Action L means to turn left the given number of degrees.
Action R means to turn right the given number of degrees.
Action F means to move forward by the given value in the direction the ship is currently facing.
The ship starts by facing east. Only the L and R actions change the direction the ship is facing.
        - x +
+   F     N 
y      W  #  E
-   R     S 
'''

f = "Day12\\12.txt"
x = 0
y = 1
position = [0, 0]
pointer = [0,0]
direction = "E"

# Need to: Determine direction of next command and apply it to the current position
# OR... should I sum up all the NSEW moves?

def change_pointer(inst):
    p = [0,0]
    if inst == "N": p = [0,1]
    elif inst == "E": p = [1,0]
    elif inst == "S": p = [0,-1]
    elif inst == "W": p = [-1,0]
    return p

#Hey turns out reading the directions is a good thing...
# def reverse_direction(d):
#     newdir = ""
#     if d == "N": newdir = "S"
#     elif d == "S": newdir = "N"
#     elif d == "E": newdir = "W"
#     elif d == "W": newdir = "E"
#     else: print("!!!!ERROR!!!!")
#     return newdir

#Turn a direction (L or R) in 90 deg increments
def rotate_ship(r, d, z):
    newdir = ""
    if r == "L":
        if d == "N": newdir = "W"
        elif d == "S": newdir = "E"
        elif d == "E": newdir = "N"
        elif d == "W": newdir = "S"
    elif r == "R":
        if d == "N": newdir = "E"
        elif d == "S": newdir = "W"
        elif d == "E": newdir = "S"
        elif d == "W": newdir = "N"
    z += -90
    #If we're rotating more than 90, we need to keep going
    if z != 0: newdir = rotate_ship(r, newdir, z)
    #I'm not happy with this, but it works I guess
    return newdir

#Open the directions file and read each line until you reach the end
for action in open(f):
    movedir = action[0]
    distance = int(action[1:])

    if movedir == "F":
        movedir = direction
    elif (movedir == "L" or movedir == "R"):
        direction = rotate_ship(movedir, direction, distance)
        movedir = direction
        distance = 0
    
    pointer = change_pointer(movedir)
    position[x] += pointer[x]*distance
    position[y] += pointer[y]*distance
    #print(position)

answer = abs(position[x]) + abs(position[y])
print("\nThe final position was %s and the answer is %s.\n" % (position, answer))

#EOF
