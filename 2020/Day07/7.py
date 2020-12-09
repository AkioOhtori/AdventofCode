goal = "shiny gold" #worst bag color ever
direct = set()
additional = set()
iterations = 0

hits = 0 #number of bags in bag set
hits_old = -1 #number of bags in previous bag set

#create initial set of bags that can contain a shiny gold one
for line in open("7.txt"):
    if goal in line:
        bags = line.find("bags")
        direct.add(line[:bags].strip())

#Loop though list until not finding any more new bags
while hits != hits_old:
    for line in open("7.txt"):
        for bag in direct:
            if bag in line:
                bags = line.find("bags")
                additional.add(line[:bags].strip()) #strip out name of bag
        direct.update(additional) #add new bags to the set
        additional.clear()

    hits_old = hits #store previous number of finds
    hits = len(direct) #update for newest finds
    iterations += 1 #count because we love data

direct.remove(goal)  #Because "shiny gold" always shows up in the results

print("\nFound " + str(len(direct)) + " bags can hold a '" + goal + "' bag in " + str(iterations) + " iterations\n")
#print the answer and we're out of here!