f = open("4.txt")
text = f.readlines()
f.close()

good = 0

valid = 0
needed = 7
x = 0
fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]

while len(text) > 0:
    ind = 0 #I don't think this ever changes
    if text[ind] == "\n":
        #validate
        if valid == needed:
            good += 1
        elif valid > needed:
            print("FML")
        valid = 0
        
    for x in range(len(fields)):
        if (text[ind].find(fields[x]) != -1):
            valid += 1
    text.pop(ind)

print("\nThere are " + str(good) + " good passports\n")
#Winner!