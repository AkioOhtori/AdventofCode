import time
start_time = time.time()

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
start = -7
    
#   100,000,000,000,000
t = 0#1000000000000000

x = start
tink = 1

while x < (len(bus)+start):
    if modsearch(x, t):
        x+=1
        tink *= bus[x][0]
    else:
        x = start
        t+= tink

print(t)
print("Execution took %s seconds\n" % (time.time() - start_time))
#EOF
