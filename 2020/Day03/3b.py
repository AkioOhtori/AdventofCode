f = open("3.txt")
good = 0
bad = 0

test = f.read()
f.close()
full = test.splitlines()

pos_x = 0
pos_y = 0
tree = []
miss = []

x_arr = [1,3,5,7,1]
y_arr = [1,1,1,1,2]

for z in range(len(x_arr)):

    inc_x = x_arr[z]
    inc_y = y_arr[z]
    print(inc_x, inc_y)

    tree.insert(z,0)
    miss.insert(z,0)

    while pos_y < len(full):
        if (full[pos_y][pos_x] == "#"):
            tree[z] += 1
        elif (full[pos_y][pos_x] == "."):
            miss[z] += 1
        else:
            print("!!!!ERROR!!!!")

        pos_x += inc_x

        if pos_x >= len(full[pos_y]):
            pos_x = pos_x - (len(full[pos_y]))
        
        pos_y += inc_y
    pos_x = 0
    pos_y = 0

print("Trees = " + str(tree))
print("Not Trees = " + str(miss))