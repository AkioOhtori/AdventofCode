# Parse input, stripping time and "x,"
floc = "Day13\\13.txt"
f = open(floc)

arrive = int(f.readline().strip())
print("\nWe arrive at %s and need to take the next bus" % arrive)
temp = (str(f.readline()))
f.close()

#NO IDEA WHY I COULDN'T GET THIS TO WORK! T_T
done = 0
while not done:
    if "x" in temp:
        xloc = temp.find("x")
        temp = temp[:xloc-1] + temp[xloc+1:]
    else:
        done = 1

bus = temp.split(",")

alltimes = list()

for n in range(len(bus)):
    time = 0
    inc = 0
    while time < arrive:
        inc +=1
        time = int((int(bus[n])*inc))
        if time >= arrive: alltimes.append([time, bus[n]])

alltimes.sort()

wait = int(alltimes[0][0]) - int(arrive)
busnum = int(alltimes[0][1])

answer = wait * busnum
print("We have to wait %s minutes for bus #%s, giving us an answer of %s!\n" % (wait, busnum, answer))
        