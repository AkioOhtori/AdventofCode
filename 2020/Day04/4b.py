f = open("4.txt")
text = f.readlines()
f.close()

good = 0

valid = 0
needed = 7
x = 0
fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
eye = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]


while len(text) > 0:
    ind = 0 #I don't think this ever changes
    if text[ind] == "\n":
        #validate
        print(valid)
        if valid == needed:
            good += 1
        elif valid > needed:
            print("FML")
        valid = 0


    temp = text[ind].find(fields[0])
    if (temp != -1):
        #birthyear = 1920-2002
        year = (text[ind][temp+4:temp+8])
        if year.isnumeric():
            year = int(year)
            if ((year >= 1920) and (year <= 2002)):
                valid += 1
    
    temp = text[ind].find(fields[1])
    if (temp != -1):
        #Issue Year
        year = (text[ind][temp+4:temp+8])
        if year.isnumeric():
            year = int(year)
            if ((year >= 2010) and (year <= 2020)):
                valid += 1

    temp = text[ind].find(fields[2])
    if (temp != -1):
        #Esp Year
        year = (text[ind][temp+4:temp+8])
        if year.isnumeric():
            year = int(year)
            if ((year >= 2020) and (year <= 2030)):
                valid += 1

    temp = text[ind].find(fields[3])
    if (temp != -1):
        #height
        height = text[ind][temp+4:temp+7].strip(" cmin")
        if (text[ind].find("in") != -1):
            if ((int(height) >= 59) and (int(height) <= 76)):
                valid += 1
        elif (text[ind].find("cm") != -1):
            if ((int(height) >= 150) and (int(height) <= 193)):
                valid += 1
    
    temp = text[ind].find(fields[4])
    if (temp != -1):
        #Hair color
        color = text[ind][temp+4:temp+11].strip()
        if ((color[0] == "#") and (color[1:6].isascii)):
            valid += 1
    
    temp = text[ind].find(fields[5])
    if (temp != -1):
        #Eye Color
        color = text[ind][temp+4:temp+7].strip()
        for x in range(len(eye)):
            if color == eye[x]:
                valid += 1
                break
    
    temp = text[ind].find(fields[6])
    if (temp != -1):
        #passport ID
        #ok this one we actually have to detect whitespace
        x = temp
        while x < len(text[ind]):
            if ((text[ind][x] == " ") or (text[ind][x] == "\n")):
                end = x
                break
            x += 1
        pid = (text[ind][temp+4:x])
        if ((pid.isnumeric) and (len(pid) == 9)):
            valid += 1



    text.pop(ind)

print("Good = " + str(good))