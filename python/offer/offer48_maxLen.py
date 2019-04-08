#coding=utf-8
#最长不含重复字符的子字符串
class Solution:

    def maxLen(self, s):
        res = 0
        cur = ""
        for v in s:
            if v not in cur:
                cur += v
                res = max(res, len(cur))
            else:
                cur += v
                cur = cur[cur.index(v)+1:]
                res = max(res, len(cur))
        return res

s = Solution()
ss = "abcdefead"
l = s.maxLen(ss)
print(l)

aa = "arabcacfr"
l = s.maxLen(aa)
print(l)

