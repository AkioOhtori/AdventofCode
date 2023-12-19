# This is gross and bad and I am sorry
# I expected this to be... easy and it was NOT haha

answer = 0
for og_line in open("Day01\\data.txt"):
    x = -1
    #for x in numbers: line = line.replace(str(x), numbers[x])
    first = ""
    second = ""
    line = og_line
    while first == "":
        x += 1
        match line[x]:
            case "o":
                if line[x:x+3] == "one": first = "1"
            case "t": #two three
                if line[x:x+3] == "two": first = "2"
                elif line[x:x+5] == "three": first = "3"
            case "f": #four five
                if line[x:x+4] == "four": first = "4"
                elif line[x:x+4] == "five": first = "5"
            case "s": #six seven
                if line[x:x+3] == "six": first = "6"
                elif line[x:x+5] == "seven": first = "7"
            case "e":
                if line[x:x+5] == "eight": first = "8"
            case "n":
                if line[x:x+4] == "nine": first = "9"
            case _:
                try:
                    first = str(int(line[x]))
                except:
                    pass
    x = 0
    while second == "":
        x += -1
        match line[x]:
            case "o":
                if line[x:x+3] == "one": second = "1"
            case "t": #two three
                if line[x:x+3] == "two": second = "2"
                elif line[x:x+5] == "three": second = "3"
            case "f": #four five
                if line[x:x+4] == "four": second = "4"
                elif line[x:x+4] == "five": second = "5"
            case "s": #six seven
                if line[x:x+3] == "six": second = "6"
                elif line[x:x+5] == "seven": second = "7"
            case "e":
                if line[x:x+5] == "eight": second = "8"
            case "n":
                if line[x:x+4] == "nine": second = "9"
            case _:
                try:
                    second = str(int(line[x]))
                except:
                    pass
    print(first + second)
    answer += int(first + second)

print(answer)

