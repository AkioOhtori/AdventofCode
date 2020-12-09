goal = "shiny gold" #still a terrible bag color
f = "7.txt"

answer = dict()  #stores the newest answers
baglist = dict()  #stores the old answers
baglist[goal] = 1
final = dict()

total = 1

while total:
    for bag in baglist: #looking for bags found in last answer
        fl = open(f)
        for line in fl:
        #for line in open(f): #go through the file line by line
            found = line.find(bag)
            if ((found < 15) and (found >= 0)): #if this line defines the bag
                #WORKING IN THE CONTEXT OF BAG COLOR "ex: purple"
                cont = line.find("contain") #pull out the bags it contains
                temp = (line[cont+8:]).strip() #parse text
                temp = temp.strip(".")
                bags = temp.split(", ") #bags this bag contains

                for x in range(len(bags)): #iterating through bags this bag contains
                    name = bags[x][2:].strip(" bag") #separate out name
                    name = name.strip(" bags")
                    qty = (bags[x][0])
                    if qty.isnumeric(): 
                        qty = int(qty)
                        qty = qty * int(baglist[bag])
                        if name in answer: answer[name] = qty + answer[name]
                        else: answer[name] = qty
                        if name in final: final[name] = qty + final[name]
                        else: final[name] = qty
                fl.close
                break
            
    baglist.clear()
    baglist = answer.copy()
    answer.clear()

    total = 0
    for z in baglist: total += baglist[z]

total = 0
for poop in final: total += final[poop]
#print(final)
print("\nBetter save your pennies, " + str(total) + " bags are required!\n")
