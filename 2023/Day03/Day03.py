f = open("Day03\\data.txt")

schematic = f.read().strip().splitlines()
numbers = []
symbols = "!@#$%^&*-+/="

score = 0
ratio = 0

for row in range(len(schematic)):
    line = schematic[row]
    x = 0
    while x < len(line):
        num = ""
        while x < len(line) and line[x].isdigit():
            num += line[x]
            x += 1
        else:
            if num != "":
                numbers.append([row, x-len(num), int(num)])
        x += 1
gear = []
for z in range(0,len(schematic)+1):
    gear.append([""]*(len(schematic[0])+1))

for number in numbers:
    x = number[1]
    y = number[0]
    n = str(number[2])

    c_x = [max([x-1,0]),min([x+len(n), len(schematic[0])-1])]
    c_y = [max([y-1,0]),min([y+1,len(schematic)-1])]

    number.append(c_y)
    number.append(c_x)

    for yy in range(c_y[0], c_y[1]+1):
        for xx in range(c_x[0], c_x[1]+1):
            if schematic[yy][xx] in symbols:
                score += int(n)
                if schematic[yy][xx] == "*":
                    gear[yy][xx] += str(n) + ","
                break
        else:
            continue
        break

print(score)

for yy in gear:
    for xx in yy:
        if xx.count(",") == 2:
            temp = xx.split(",")
            ratio += int(temp[0])*int(temp[1])
print(ratio)
