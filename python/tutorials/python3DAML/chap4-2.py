def guess(x, y):
    if x < y:
        print('bigger')
    elif x > y:
        print('smaller')
    else:
        print('correct')
        return True
    return False

def main():
    y = input()#target
    x = input()#guess
    while guess(x, y) == False:
        x = input()

if __name__ == '__main__':
    main()
