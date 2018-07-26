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

class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        return ' '.join([d[::-1] for d in s.split(' ') if d!=''])
if __name__ == '__main__':
    s = Solution()
    print(s.reverseWords('the sky is blue'))#eht yks si eulb
    print(s.reverseWords('Let\'s take LeetCode contest'))#s'teL ekat edoCteeL tsetnoc
    # print(reverseWords("the sky is blue"))
    # print(reverseWords("    the sky is blue   "))
    # print(reverseWords("s     a"))
