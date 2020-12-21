import numpy as np

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

matrix = []
oldmatrix = []

#final form = matrix[z][y][x]
for y in range(len(temp)):
    temp2 = []
    for x in range(len(temp[y])):
        if temp[y][x] == "#": temp2.append(active)
        else: temp2.append(inactive)
    oldmatrix.append(temp2.copy())

#OK.  OK.  OK.... OK.
#What are we doing!?

#Check if an element is active or inactive
#If active, check adjacent until we see more than 3
#If inactive, check adjacent until we see more than 3

def checkadjacent_2d(mmm,cx,cy):
    #Recieves a 2D matrix and x,y coords and returns adjacent active cells
    ctr = 0
    starty = -1
    startx = -1
    #this is a messy way to deal with negative indexes but meh
    if cx == 0: startx = 0
    if cy == 0: starty = 0
    for y in range(starty,2):
        for x in range(startx,2):
            try:
                if mmm[y+cy][x+cx] == active: ctr+=1
            except: pass
            #if ctr > inactive_minmax+1: return 99
    #if m[cy][cx] == active: return ctr-1
    #else: return ctr
    return ctr

def checkadjacent_3d(m, cx, cy, cz):
    ctr = 0
    startz = -1
    if cz == 0: startz = 0
    for z in range(startz, 2):
        ctr += checkadjacent_2d(m[z], cx, cy)
    if m[cz][cy][cx] == active: return ctr-1
    else: return ctr

#takes a 2D! matrix and returns a matrix with 0s on peremiter
def growmatrix_xy(m):
    for y in range(len(m)):
        m[y].append(0)
        m[y].insert(0, 0)
    yy = [0] * len(m[0])
    m.append(yy.copy())
    m.insert(0,yy.copy())
    return m

#takes a size and returns a SQUARE 0 matrix of same size
def zeromatrix(s):
    x = [0]*int(s)
    n = []
    while s > 0:
        n.append(x.copy())
        s += -1
    return n

#takes in a 2 or 3D matrix and returns zero padded matrix
def growmatrix_z(m):
    try:
        l = int(len(m[0][0]))
        z = zeromatrix(l).copy()
        m.append(z.copy())
        m.insert(0, z.copy())
    except:
        n = []
        z = zeromatrix(int(len(m))).copy()
        z = list(map(lambda a: a.copy(), z))
        zz = [a.copy() for a in z]        
        print(np.matrix(z))
        n.append(z.copy())
        n.append(m.copy())
        n.append(zz.copy())
        m = n.copy()

    return m

#takes in a 3D matrix and pads it on all sides with 0
def growmatrix(m):
    for z in range(len(m)):
        m[z] = growmatrix_xy(m[z]).copy()
    m = growmatrix_z(m.copy()).copy()
    return m

testmode = 0

def changematrix_3d(m):
    nxyz = []
    for z in range(len(m)):
        nxy = []
        for y in range(len(m[z])):
            nx = []
            for x in range(len(m[z][y])):
                a = checkadjacent_3d(m,x,y,z)
                if testmode == 1:
                    nx.append(a)
                else:
                    if a == 3: nx.append(1)
                    elif a == 2 and m[z][y][x] == active: nx.append(1)
                    else: nx.append(0)
            nxy.append(nx.copy())
        nxyz.append(nxy.copy())
    return nxyz

def count_the_ones(m):
    a = 0
    for z in m:
        for y in z:
            a += y.count(1)
    return a

print(oldmatrix)

#Prep because feeding 3D functions the starter matrix kept breaking
matrix = growmatrix_xy(oldmatrix).copy()
oldmatrix = growmatrix_z(matrix).copy()
oldmatrix[0][2][2] = 1
print(np.matrix(oldmatrix[0]))
print(np.matrix(oldmatrix[1]))
print(np.matrix(oldmatrix[2]))
print("Matrix for cycle %s is %sx, %sy, %sz and has %s 1s" % (1, len(oldmatrix[1][1]), len(oldmatrix[1]), len(oldmatrix), count_the_ones(oldmatrix)))
cycle_end = 6
cycle_count = 1
#print(np.matrix(A))
print((checkadjacent_3d(oldmatrix, 2, 2, 1)))



while cycle_count < cycle_end:
    matrix = changematrix_3d(oldmatrix).copy()
    oldmatrix = growmatrix(matrix).copy()
    #oldmatrix = matrix.copy()
    cycle_count += 1
    print("Matrix for cycle %s is %sx, %sy, %sz and has %s 1s" % (cycle_count, len(oldmatrix[0][0]), len(oldmatrix[0]), len(oldmatrix), count_the_ones(oldmatrix)))
#print(matrix)

answer = 0

print(answer)