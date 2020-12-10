# Adapter rating = +0j, - 1-3j
# Build in adatper = highest + 3j
# Outlet = 0j
# How can you use all the adapters?  What are the voltage differences?

# Uhh this seems super simple?

f = open("10.txt")
test = f.read().splitlines()
f.close()
adapters = [0] #0 for outlet

differences = dict()
differences[1] = 0
differences[2] = 0
differences[3] = 0

for a in range(len(test)): adapters.append(int(test[a].strip()))

adapters.sort() #sort smallest to largest
adapters.append(adapters[len(adapters)-1] + 3) #append internal adapter

#Iterate though sorted list and store the difference type in the dict
for x in range(len(adapters)-1): differences[adapters[x+1] - adapters[x]] += 1

#Per problem statement, answer is 1j diffs multiplied by 3j diffs
answer = differences[1] * differences[3]
print("\nThe answer is " + str(answer))

#EOF
