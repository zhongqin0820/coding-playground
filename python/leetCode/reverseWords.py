def reverseWords(s):
    """
    :type s: str
    :rtype: str
    """
    data = [v for v in s.split(' ') if v!='']
    i, j = 0, len(data)-1
    if i>j:return ''
    while i < j:
        data[i], data[j] = data[j], data[i]
        i, j = i+1, j-1
    return ' '.join(data)

if __name__ == '__main__':
    print(reverseWords("the sky is blue"))
    print(reverseWords("    the sky is blue   "))
    print(reverseWords("s     a"))
