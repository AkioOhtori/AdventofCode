test = "one1two2three3"


numbers = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9"
}

for x in numbers:
    test = test.replace(str(x), numbers[x])


test = "one1two2three3"

test = test[0:3].replace("one","1") + test[3:]

print(test[-4:-1])