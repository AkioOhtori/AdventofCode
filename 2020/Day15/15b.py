import time
start_time = time.time()

#Input processing and setup. Yes, there is probably a better way to do this
floc = "Day15\\15.txt"
f = open(floc)
i = (f.readline().strip().split(","))
f.close()

#set up variables
game = dict()
last = 0
turn = 1
say = 0
end = 30000000#2020

#Seed with initial values
for n in i:
    game[int(n)] = [0, turn]
    last = int(n)
    turn += 1

#Play the game
while turn <= end:
    #check if this has been said only once
    if game[last][0] == 0: say = (turn - game[last][1] -1)
    else: say = game[last][1] - game[last][0]
    last = say
    if say not in game: game[say] = [0,0] #int if it doesn't already exist

    del game[say][0] #remove oldest turn
    game[say].append(turn) #add in this one
    turn += 1 #next!

print("\nThe %sth number said was %s! That took %s seconds to compute.\n" % (end,say, (("%.1f" % (time.time() - start_time)))))