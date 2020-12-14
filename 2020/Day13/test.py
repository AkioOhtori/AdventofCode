
floc = "Day13\\13.txt"
f = open(floc)
x = f.readline().strip()
bus = [int(bus) for bus in f.readline().split(',') if bus != 'x']
print(bus)