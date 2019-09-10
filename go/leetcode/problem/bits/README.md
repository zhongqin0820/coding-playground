# Date
创建：2019-09-10

# 相关链接
- [CS-Notes/notes/Leetcode 题解 - 位运算.md](https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E4%BD%8D%E8%BF%90%E7%AE%97.md#1-%E7%BB%9F%E8%AE%A1%E4%B8%A4%E4%B8%AA%E6%95%B0%E7%9A%84%E4%BA%8C%E8%BF%9B%E5%88%B6%E8%A1%A8%E7%A4%BA%E6%9C%89%E5%A4%9A%E5%B0%91%E4%BD%8D%E4%B8%8D%E5%90%8C)

# Go中相关的包
- math/bits

## 主要函数说明
- `Add`：返回结果带进位
- `Div`：返回商，余数
- `Len`：统计返回表示某数需要的位数
- `Mul`：
- `Reverse`：翻转数
- `RotateLeft`：将数向左循环移动k位；当k为负时表示向右循环移动
- `Sub`：带借位的减法
- `OnesCount`：统计返回含1的位数
- `LeadingZeros`：统计返回左起遇到1前0的个数
- `TrailingZeros`：统计返回右起遇到1前0的个数

## 基本操作
- `i>>1 == i/2`
- `i&1 == i%2`
- `i&(-i)`：找到右起第一个1所对应的值
- `i&(i-1)`：去除右起第一个1所对应的值
