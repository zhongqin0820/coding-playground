# Date
文档创建：2019-09-13

# 相关链接
- [CS-Notes/notes/Leetcode 题解 - 排序.md](https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E6%8E%92%E5%BA%8F.md)

# 涉及
## 快速排序
- 重点掌握切分部分的应用
- `"sort"`
    - `sort.Sort(sort.Interface)`: 自定义的`sort.Interface`需要实现三个方法签名
    - 自带基础数据结构切片排序方法:
    - `sort.Ints`: 升序
    - `sort.Floats`: 升序
    - `sort.Strings`: 字典序升序

## 堆排序
- `"container/heap"`: 需要实现五个方法签名

## 桶排序
- 统计频率，对频率排序，取top k
