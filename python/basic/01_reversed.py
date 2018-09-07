# reversed output a iterator
a = [1, 2, 3]
print(a)
print(type(a))
b = (1,2,3,2,4)
print(b)
print(type(b))
c = list({4,5,6})
print(c)
d = tuple({4,5,6})
print(d)
e = set(c)
f = dict({'a':1, 'b':2})
print(f)
aa = reversed(a)
print(aa)
for i in aa:
    print(i)
