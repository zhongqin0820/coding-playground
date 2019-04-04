# -*- coding:utf-8 -*-
import heapq

class Solution:
    def GetLeastNumbers_Solution(self, t, k):
        # write code here
        if not t or len(t)<k:
            return []
        return self.heap(t,k)
        # return self.helper(t)[:k]

    def helper(self, t):
        if len(t)<2:
            return t[:]
        left = self.helper([i for i in t[1:] if i <= t[0]])
        right = self.helper([i for i in t[1:] if i > t[0]])
        return left+[t[0]]+right

    def heap(self, t, k):
        if not t or len(t)<k:
            return []
        res = []
        for v in t:
            heapq.heappush(res, -v) if len(res)<k else heapq.heappushpop(res, -v)
        return sorted(list(map(lambda v: -v, res)))[:k]

s = Solution()
res = s.GetLeastNumbers_Solution([3,5,4,6,8,9,4],4)
print(res)
