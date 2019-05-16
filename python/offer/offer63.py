# coding:utf-8
# best time to buy and sell stock
class Solution:
    
    # basic version
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 1 or prices ==  None:
            return 0
        min_, max_ = prices[0], 0
        for v in prices:
            min_ = min(min_, v)
            max_ = max(max_, v-min_)
        return max_

    # problem 309: with cooldown version
    def maxProfit_309(self, prices: List[int]) -> int:
        if len(prices) < 2 or prices == None:
            return 0
        n = len(prices)
        buy,sell = [0] * n,[0] * n
        buy[0] = -prices[0]
        buy[1],sell[1] = max(buy[0],-prices[1]),max(0,buy[0]+prices[1])
        for i in range(2,n):
            buy[i] = max(buy[i-1],sell[i-2]-prices[i])
            sell[i] = max(sell[i-1],buy[i-1]+prices[i])
        return sell[n-1]
