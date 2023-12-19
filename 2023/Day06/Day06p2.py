# 1 ms hold = 1mm/ms gain
race = []

f = open("Day06\\data.txt")
data = f.read().strip().splitlines()
for x in range(len(data)):
    data[x] = data[x][data[x].find(":")+1:].replace(" ", "")
race.append(int(data[0]))
race.append(int(data[1]))


ht = 1
wins = 0
while ht <= race[0]:
    d = ht * (race[0] - ht)
    if d > race[1]: wins += 1
    ht += 1

print(wins)
# This is a little slow but not bad
# I suspect the right way would be to find the first and last win and subtract