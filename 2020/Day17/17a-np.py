import numpy as np
import time

f = open("Day17\\17s.txt")
test = f.read()
f.close()
temp = test.splitlines()
print(temp)

active = 1
inactive = 0

active_min = 2
active_max = 3
inactive_minmax = 3

temp2 = []#np.array([int()])
#final form = matrix[z][y][x]
for y in range(len(temp)):

    for x in range(len(temp[y])):
        if temp[y][x] == "#": temp2.append(active)
        else: temp2.append(inactive)
matrix_2d = np.array(temp2, copy=True)
matrix_2d = matrix_2d.reshape(len(temp[0]),len(temp)).copy()
print(np.array(matrix_2d))

def growmatrix_xy(m_2d):
    m_2d = np.insert(m_2d, 0, 0, axis=0)
    m_2d = np.insert(m_2d, len(m_2d), 0, axis=0)
    m_2d = np.insert(m_2d, 0, 0, axis=1)
    m_2d = np.insert(m_2d, len(m_2d[0]), 0, axis=1)
    return m_2d

def searchmatrix():
    return 0

def countaround(x,y,z,m):
    #poop = np.where(m[z-1:z+2, y-1:y+2, x-1:x+2])
    #print(len(poop[0])-1)
    if m[z,y,x]: return (np.count_nonzero(m[z-1:z+2, y-1:y+2, x-1:x+2])-1)
    else: return (np.count_nonzero(m[z-1:z+2, y-1:y+2, x-1:x+2]))


while matrix_2d.shape[0] < 17:
    matrix_2d = growmatrix_xy(matrix_2d)

cube = (np.zeros((15,17,17), dtype=int))
cube[8] = matrix_2d.copy()

#print(np.dstack((cube,matrix_2d,zero2)))
# print(matrix_2d[..., np.newaxis])
# b = np.reshape(matrix_2d, (-1,5,5))
# b = np.insert(b, 0, cube, axis = 0)
# b = np.insert(b, 2, zero2, axis = 0)

# print(b.shape)

# growmatrix_z(matrix_2d)

# hits = (np.where(cube == 1))
# print(np.array(hits))
# for n in range(len(hits[0])):
#     x = hits[2][n]
#     y = hits[1][n]
#     z = hits[0][n]
#     countaround(x, y, z, cube)

# limited_search = np.array(np.where(cube[5:7]))
# print(len(limited_search[0]))
# print(cube[5:8,5:8,6:9])
iteration = 0
while iteration < 6:
    cube_old = cube.copy()
    for z in range(1,14):
        for y in range(1,16):
            for x in range(1,16):
                c = countaround(x,y,z,cube_old)
                # if c == 3: print(x,y,z)
                if c == 3: cube[z,y,x] = 1
                elif c != 2 and cube_old[z,y,x] == 1: cube[z,y,x] = 0
    print("On iteration %s there were %s 1's found!" % (iteration, np.count_nonzero(cube)))
    iteration += 1
print(np.count_nonzero(cube))