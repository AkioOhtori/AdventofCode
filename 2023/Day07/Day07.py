
from collections import Counter

total = 0
winnings = 0
hands = {
    "high": [],
    "pair": [],
    "twop": [],
    "three": [],
    "full": [],
    "four": [],
    "five": [],
}

points = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2,
}

for line in open("Day07\\data.txt"):
    line = line.strip().split(" ")
    bet = int(line[1])

    type = list(Counter(line[0]).values()) #extract only the values for ease of processing
    #hand = ''.join((sorted(line[0],key=line[0].count,reverse=True))) #ordered hand
    hand = []
    for l in line[0]: hand.append(points[l])

    type.sort(reverse=True)

    match type[0]:
        case 5:
            hands["five"].append([hand.copy(),hand,bet])
        case 4:
            hands["four"].append([hand.copy(),hand,bet])
        case 3:
            if type[1] == 2:
                hands["full"].append([hand.copy(),hand,bet])
            else:
                hands["three"].append([hand.copy(),hand,bet])
        case 2:
            if type[1] == 2:
                hands["twop"].append([hand.copy(),hand,bet])
            else:
                hands["pair"].append([hand.copy(),hand,bet])
        case 1:
            hands["high"].append([hand.copy(),hand,bet])
        case _:
            print("Error ", type)

# ok now we should be ready to start processing bets... oh god I forgot about those crap

for t in hands:
    hands[t].sort()
    for x in range(len(hands[t])):
        total += 1
        winnings += hands[t][x][2] * total
print(winnings)