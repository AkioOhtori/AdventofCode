f = open("3.txt")
good = 0
bad = 0

test = f.read()
f.close()
full = test.splitlines()

pos_x = 0
pos_y = 0
tree = 0
miss = 0

inc_x = 3
inc_y = 1

for y in range(len(full)):

    if (full[y][pos_x] == "#"):
        tree += 1
    elif (full[y][pos_x] == "."):
        miss += 1
    else:
        print("!!!!ERROR!!!!")

    pos_x += inc_x
    #pos_y += inc_y

    if pos_x >= len(full[y]):
        pos_x = pos_x - (len(full[y]))

print("Trees = " + str(tree))
print("Not Trees = " + str(miss))