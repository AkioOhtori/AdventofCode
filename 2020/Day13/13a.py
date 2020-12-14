# Parse input, stripping time and "x,"
floc = "Day13\\13.txt"
f = open(floc)

arrive = int(f.readline().strip())
print("\nWe arrive at %s and need to take the next bus" % arrive)
bus = [int(bus) for bus in f.readline().split(',') if bus != 'x']
#Thanks to @duplico for this neat line of code.  My version was gross.

alltimes = list()

for n in range(len(bus)):
    time = 0
    inc = 0
    while time < arrive:
        inc +=1
        time = int((int(bus[n])*inc))
        #Only store valid times
        if time >= arrive: alltimes.append([time, bus[n]])

alltimes.sort()
#[x][0] = time; [x][1] = bus number
wait = int(alltimes[0][0]) - int(arrive)
busnum = int(alltimes[0][1])

answer = wait * busnum #per problem statement
print("We have to wait %s minutes for bus #%s, giving us an answer of %s!\n" % (wait, busnum, answer))
        