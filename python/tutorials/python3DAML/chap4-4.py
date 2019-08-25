# from chap4-1 import fib
def fib_smaller_than_(n):
    a = list()
    for i in range(0,n):
        if i == 0 or i == 1:
            a.append(1)
        else:
            v = a[i-2] + a[i-1]
            if v > n:
                break
            else:
                a.append(v)
    return a


def sum_to_n(n):
    sum = 0
    for i in range(1,n+1):
        sum += i
    return sum

def main():
    n = input()
    print(fib_smaller_than_(n))
    print(sum_to_n(n))

if __name__ == '__main__':
    main()
