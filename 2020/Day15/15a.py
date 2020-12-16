#Input processing and setup. Yes, there is probably a better way to do this
floc = "Day15\\15.txt"
f = open(floc)
i = (f.readline().strip().split(","))
f.close()

game = dict()
last = 0
turn = 1
say = 0
end = 2020
temp = 0

for n in i:
    game[int(n)] = [0, turn]
    last = int(n)
    turn += 1
    print(game[last])
print(game)

while turn <= end:
    if game[last][0] == 0: say = (turn - game[last][1] -1)
    else: say = game[last][1] - game[last][0]
    last = say
    if say not in game: game[say] = [0,0] #int if it doesn't already exist

    del game[say][0] #remove oldest turn
    game[say].append(turn) #add in this one
    #print(say)
    turn += 1 #next!

print("\nThe %sth number said was %s!\n" % (end,say))