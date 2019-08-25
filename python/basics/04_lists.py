a = [1, -1, 3, 2, 5, 4]
# append
a.append(6)
print('a:', a)
# sorted
b = sorted(a)
print('b:', b)
print('a:', a)
# a.sort()
# print('a:', a)
# slice
# skip step=2
c = a[::2]
print(c)
# sawp a[i], a[i+1]
c = a[::-1]
print(c)
# iterate an array
s = ', '.join([str(v) for v in a])
print('s:', s)
# list comprehensions pro
l = ', '.join([str(x+y) for x in a for y in b])
print(l)
# iterator v/s generator
g = (x*x for x in a)
print(next(g))
for v in g:
    print(v)
# self define a generator function
def genX(a):
    for v in a:
        print('yield: ', v)
        yield(v)
x = genX(a)
while True:
    try:
        next(x)
    except StopIteration as e:
        print('error: ', e.value)
        break
