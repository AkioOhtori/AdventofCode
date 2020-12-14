
'''
Action N means to move the waypoint north by the given value.
Action S means to move the waypoint south by the given value.
Action E means to move the waypoint east by the given value.
Action W means to move the waypoint west by the given value.
Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
Action F means to move forward to the waypoint a number of times equal to the given value
        - x +
+   F     N 
y      W  #  E
-   R     S 
'''

# Part 2 tl;dr is NSEW moves the WAYPOINT
#                 F moves the SHIP
#                 LR rotates the WAYPOINT


f = "Day12\\12.txt"
x = 0
y = 1
position = [0, 0]
waypoint = [10, 1]

#Rotate the waypoint coordinants (L or R) in 90 deg increments
def rotate_waypoint(w, r, z):
    if r == "L":
        newwaypoint = [-w[y], w[x]]
    else: newwaypoint = [w[y], -w[x]]
    z += -90
    if z != 0: newwaypoint = rotate_waypoint(newwaypoint, r, z)
    return newwaypoint

# Waypoint must move in the perscribed direction
def move_waypoint(w, inst, z):
    if inst == "N": w[y] += z
    elif inst == "E": w[x] += z
    elif inst == "S": w[y] += -z
    elif inst == "W": w[x] += -z
    return w

# Move the ship z number of times towards the waypoint
def move_ship(w, p, z):
    newpositon = [(p[x]+(w[x]*z)), (p[y]+(w[y]*z))]
    return newpositon

#Open the directions file and read each line until you reach the end
for action in open(f):
    movedir = action[0]
    distance = int(action[1:])

    if (movedir == "L" or movedir == "R"):
        waypoint = rotate_waypoint(waypoint, movedir, distance)
    elif (movedir == "F"):
        position = move_ship(waypoint, position, distance)
    else: #Instruction is NSEW
        waypoint = move_waypoint(waypoint, movedir, distance)
    #print(position)

answer = abs(position[x]) + abs(position[y])
print("\nThe final position was %s and the answer is %s.\n" % (position, answer))

#EOF
