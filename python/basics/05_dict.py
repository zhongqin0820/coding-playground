# definition
d1 = dict([('a', 1), ('b', 2)])
print(d1)
d2 = {'a':1, 'b':2}
print(d2)
d3 = dict(a=1, b=2)
print(d3)
# keyError
try:
    print(d3[1])
except KeyError as e:
    print(type(e))
    d3[e] = 3
    print(e in d3)
    print(d3[e])
d3.clear()
print(d3)
# items, keys, values
print(d2.items())
print(d2.keys())
print(d2.values())
# get and pop with default value
print(d2.get('a', 4))
print(d2.get('c', 4))
print(d2.pop('a', -1))
print(d2.pop('a', -1))
# pop items at the end
print(d1.popitem())
print('d1', d1)
# update items
print(d2.update([('b', 4), ('e', 9)]))
print(d2)
