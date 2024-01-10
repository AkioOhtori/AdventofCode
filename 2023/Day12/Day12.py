for line in open("Day12\\sample.txt"):
    non = (line.strip().split())
    # non[1] = non[1].split(",")
    non[1] = list(map(int,non[1].strip().split(",")))
    print(non)

    #ok so what can we tell from what we have?
    #   number of groups based on number of... groups haha
    #   total minimum spaces from group size plus required whitespace
    #   ??
    groups = len(non[1])
    white = (groups -1)
    space = sum(non[1]) + white
    available = len(non[0])
    wild = non[0].count("?")
    damaged = non[0].count("#")
    operating = non[0].count(".")
    print(groups, white, space, available)
          
    print("Damaged: %s; Operating: %s; Wild: %s\n" % (damaged, operating, wild))

    # for sample, answers are 1, 4, 1, 1, 4, 10
    # JUST looking at the grouping we can get close, but not there yet

