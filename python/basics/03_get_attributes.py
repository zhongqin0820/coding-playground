# define a student class
class student:
    def __init__(self, age, name):
        self.age = age
        self.name = name
        return

    def get_age(self):
        """
        :rtype: int
        """
        return self.age
    def get_name(self):
        """
        :rtype: string
        """
        return self.name

# use dir() and type() built-in function to get attributes of a class 
if __name__ == '__main__':
    me = student(21, 'me')
    print([i for i in dir(me) if i[0]!='_'])
    print(type(me))
    print(isinstance(me, student))
    # be sure to assign the parameter type string
    print(hasattr(me, 'age'))
    print(hasattr(me, 'noexistedAttributes')) # return False
    print(getattr(me, 'name'))
    # the code below throw AttributeError, we need to assign an default return flag or use try...catch..
    # print(getattr(me, 'noexistedAttributes')) 
    # Method 1:
    try:
        print(getattr(me, 'noexistedAttributes')) 
    except AttributeError as e:
        print(e)
    finally:
        print('finally...')
    # Method 2:
    print(getattr(me, 'noexistedAttributes', 'error, not found!'))
