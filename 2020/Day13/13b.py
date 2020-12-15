# Parse input, stripping time and "x,"
floc = "Day13\\13.txt"
f = open(floc)

arrive = int(f.readline().strip())
buses = f.readline().strip().split(",")
f.close()

test = []
z = 0
for bus in buses:
    if bus != "x":
        test.append([int(bus), int(z)])
    z += 1
print(test)
bus = test

def modsearch(n, t):
    if ((t+bus[n][1])%(bus[n][0]) == 0): return 1
    else: return 0

t = 100000000000000
x = 0
while x < len(bus):
    if modsearch(x, t):
        x+=1
    else:
        x = 0
        t+=1

print(t)
