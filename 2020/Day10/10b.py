# Adapter rating = +0j, - 1-3j
# Build in adatper = highest + 3j
# Outlet = 0j
name = "2020\\Day10\\10.txt"
f = open(name)
test = f.read().splitlines()
f.close()
adapters = [0] #0 for outlet
arragements = 0
iterations = 0

for a in range(len(test)): adapters.append(int(test[a].strip()))

adapters.sort() #sort smallest to largest
adapters.append(adapters[len(adapters)-1] + 3) #append internal adapter
end = len(adapters)-1

def runout(start, ans):
    global iterations
    if start == end:
        return (ans+1)
    for n in range(1,4):
        if (start + n) <= end:
            if (adapters[start + n] - adapters[start]) <= 3:
                ans = runout((start+n), ans)
    iterations +=1
    #print(iterations)
    return ans
answer = 0
answer = runout(0,0)

print(answer)

#EOF
