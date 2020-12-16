

with open("Day13\\13.txt") as F:
    F.readline() # ignore the first line of the input
    buses = F.readline().strip().split(',')

# create pairs of (divisor, remainder) for every available bus
buses = [(int(buses[i]), (int(buses[i]) - i) % int(buses[i]))
    for i in range(len(buses)) if buses[i] != 'x']
print(buses)
result = 0
increment = 1

for bus in buses:
    while result % bus[0] != bus[1]:
        print("Tried %s %% bus %s but which was %s not %s. Increasing by %s." % (result, bus[0], (result % bus[0]), bus[1], increment))
        result += increment
    increment *= bus[0]
    print("Success! New increment is %s." % increment)

print(result)