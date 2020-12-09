
f = open("1.txt")
#print(f.read())

test = f.read()
a = test.splitlines()
b = a
c = a

for x in range(len(a)):
    for y in range(len(b)):
        for z in range(len(c)):
            math = int(a[x]) + int(b[y]) + int(c[z])
            if ((math) == 2020):
                ans = int(a[x])*int(b[y])*int(c[z])
                print(ans)
                break
print("done")
