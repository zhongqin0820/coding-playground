from collections import *
# namedtuple
# define an object simply
founder = namedtuple('founder', ['url', 'name', 'founder'])
websites = [
	('c', 'c', 'c'),
	('a', 'a', 'a'),
	('b', 'b', 'b')
]
for w in websites:
    f = founder._make(w)
    print(f)
# deque
# use if the operation of delete and append from head is frequent
# import sys
# import time
# fancy_load = deque('>-----')

# while True:
#     print('\r%s'%''.join(fancy_load))
#     fancy_load.rotate(1)
#     sys.stdout.flush()
#     time.sleep(0.08)
d = deque([1,2,3])
print(d)
# print(type(d))
# for ed in d:
#     print(ed)
d.append('a')
d.appendleft('b')
print(d)
ed = d.pop()
print(ed)
ed = d.popleft()
print(ed)
# defaultdict
dd = defaultdict(lambda: 'N/A')
dd['a'] = 1
print(dd['a'])
print(dd['b'])
# OrderedDict
od = OrderedDict({'a':1, 'b':2, 'c':3})
ud = dict({'a':1, 'b':2, 'c':3})
print(od)
print(ud)
print(list(od.keys()))
# Counter
# Counter is the sub set of dict, the same as a counting set
c = Counter()
for ch in 'this is a sentence':
    c[ch] = c[ch] + 1
print(c)
