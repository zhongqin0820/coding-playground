# initialization
s = set([1,1,1,1,2,1,2])
u = set([1,2,3,4])
print(s)
# create
s.add(8)
u.add(9)
# t here shares the same memory with s
t = s
print(s)
# retrive
print(1 in s)
# update
t.update(u)
# delete
s.remove(1)
print('s:',s)
print('t:',t)
# math operation
v = s | u
print('v:',v)
v = s & u
print('v:',v)
v = s - u
print('v:',v)
v = s ^ u
print('v:',v)
