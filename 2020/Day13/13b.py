# Parse input, stripping time and "x,"
floc = "Day13\\13s.txt"
f = open(floc)

arrive = int(f.readline().strip())

buses = f.readline().strip().split(",")

alltimes = list()
print(buses)

test = []
z = 0
for bus in buses:
    if bus != "x":
        test.append([int(bus), int(z)])
    z += 1
print(test)
bus = test

print(bus[-3])

def search(b, n):
    #for p in range(len(test)):
    if b > 1:#len(bus):
        print("FINISHED")
        print(b)
    a1 = 0
    a2 = 0
    p = 1
    while a2 <= a1:
        a1 = bus[b][0]*n + bus[b][1]
        a2 = bus[b+1][0]*p + bus[b+1][1]

        if (a1 == a2):
            #found a hit, move to next sequence
#            print("HIT")
            search(b+1, p)
        else: p += 1



t=1
while t < 1068797:
    search(-3, t)
    t += 1

print(t)
