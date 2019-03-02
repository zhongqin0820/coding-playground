'''
fibnoci: 1 1 2 3 5 8..
a_i = 1 (i=0,1)
a_i = a_{i-2} + a_{i-1} (i>1)
'''
import logging


def fib(i):
    a = list()

    for idx in range(0,i):
        if idx == 0 or idx == 1:
            a.append(1)
        else:
            a.append(a[idx-2]+a[idx-1])
    return a

def main():
    print('enter an intenger:')
    i = raw_input()

    a = fib(i)

    for item in a:
        print(item + ",")

if __name__ == '__main__':
    main()
