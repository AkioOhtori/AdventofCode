import numpy as np
import time

f = open("Day17\\17.txt")
test = f.read()
f.close()
temp = test.splitlines()

active = 1
inactive = 0

cycles = 6

temp2 = []
#Read input data from file and slice into a np array
for y in range(len(temp)):

    for x in range(len(temp[y])):
        if temp[y][x] == "#": temp2.append(active)
        else: temp2.append(inactive)
matrix_2d = np.array(temp2, copy=True)
matrix_2d = matrix_2d.reshape(len(temp[0]),len(temp)).copy()
print("Starting input: ")
print(np.array(matrix_2d))

#Determine the size of the starting data
startx = len(matrix_2d[0])
starty = len(matrix_2d)

#Scope the zero matrix to be 1x larger than needed on every side
sizex = startx + 2 * (cycles+1)
sizey = starty + 2 * (cycles+1)
sizez = 1 + 2 * (cycles+1)

#add zeros to a 2D matrix on all sides - probably a better way to do this
def growmatrix_xy(m_2d):
    m_2d = np.insert(m_2d, 0, 0, axis=0)
    m_2d = np.insert(m_2d, len(m_2d), 0, axis=0)
    m_2d = np.insert(m_2d, 0, 0, axis=1)
    m_2d = np.insert(m_2d, len(m_2d[0]), 0, axis=1)
    return m_2d

def countaround(x,y,z,m):
    c = (np.count_nonzero(m[z-1:z+2, y-1:y+2, x-1:x+2]))
    if m[z,y,x]: return c-1 #if we've counted ourselves, remove
    else: return c

#Grow the starting data to the size needed for easy insertion (heyo)
while matrix_2d.shape[0] < sizex:
    matrix_2d = growmatrix_xy(matrix_2d)

#Create a cube of 0s the size we determined earlier
cube = (np.zeros((sizez,sizey,sizex), dtype=int))
cube[int(len(cube)/2)] = matrix_2d.copy() #insert starting data

iteration = 1
#Exectue puzzle rules on every element until we're done
while iteration <= cycles:
    cube_old = cube.copy()
    for z in range(1,sizez-1):
        for y in range(1,sizey-1):
            for x in range(1,sizex-1):
                c = countaround(x,y,z,cube_old)
                if c == 3: cube[z,y,x] = 1
                elif c != 2 and cube_old[z,y,x] == 1: cube[z,y,x] = 0
    print("On iteration %s there were %s 1's found!" % (iteration, np.count_nonzero(cube)))
    iteration += 1
#EOF