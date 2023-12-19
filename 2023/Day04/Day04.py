from math import floor
score = 0
total = 0
matches = []

for line in open("Day04\\data.txt"):
    data = line[line.find(":")+2:]
    ldata = data[:data.find("|")-1]
    rdata = data[data.find("|")+2:]
    winners = ldata.replace("  "," ").split(" ")
    numbers = rdata.strip().replace("  "," ").split(" ")
    
    points = -1

    for n in winners:
        if n in numbers:
            points += 1
    score += floor(pow(2,points))
    matches.append(points+1)

print(score)
bonus = [1]*len(matches)

for z in range(len(matches)):
    x = bonus[z]
    total += bonus[z]
    for n in range(matches[z]):
        try:
            bonus[n+z+1] += x
        except:
            pass

print(total)