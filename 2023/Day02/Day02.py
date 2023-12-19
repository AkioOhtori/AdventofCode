game = 0

r_max = 12
g_max = 13
b_max = 14

valid = 0
power = 0

for line in open("Day02\\data.txt"):
    game += 1
    rgb = [[],[],[]]
    draws = line.strip()
    draws = line[draws.find(":")+2:].strip()
    draw = draws.split("; ")
    for hand in draw:
        cubes = {
            "red": 0,
            "blue": 0,
            "green": 0
        }

        hand = hand.replace(",","")
        values = hand.split(" ")

        for x in range(1, len(values), 2):
            cubes[values[x]] += int(values[x-1])
        rgb[0].append(cubes["red"])
        rgb[1].append(cubes["green"])
        rgb[2].append(cubes["blue"])
    if (max(rgb[0])) <= r_max and (max(rgb[1])) <= g_max and (max(rgb[2])) <= b_max:
        valid += game
    p = max(rgb[0])*max(rgb[1])*max(rgb[2])
    power += p
print(valid) # part 1
print(power) # part 2