# -*- coding:utf-8 -*-
import heapq
class Solution:
    def __init__(self):
        self.heaps = [], []#小根堆(堆顶最大,取反即最小), 大根堆(堆顶最大,堆底最小)
    def Insert(self, num):
        # write code here
        left, right = self.heaps
        heapq.heappush(left, -heapq.heappushpop(right,num))#往大根堆放,放不下,就往小根堆放,记得要取反
        if len(left)>len(right):#调整堆,保证左边小于右边
            heapq.heappush(right,-heapq.heappop(left))#将堆顶弹出给大根堆
        print('left', left, 'right', right)
        
    def GetMedian(self,s):
        # write code here
        left, right = self.heaps
        if len(right)>len(left):
            return float(right[0])
        return (right[0]-left[0])/2.0#最大堆的最小+最小堆的最大(元素本身为负),取二者平均
s = Solution()
s.Insert(3)
s.Insert(6)
s.Insert(1)
s.Insert(8)
s.Insert(10)
s.Insert(2)
print(s.GetMedian(""))
