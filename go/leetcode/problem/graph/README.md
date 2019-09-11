# Date
创建：2019-09-11

# 学习资料
- [CS-Notes/notes/Leetcode 题解 - 图.md](https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%9B%BE.md)
- [CS-Notes/notes/算法 - 并查集.md](https://github.com/CyC2018/CS-Notes/blob/master/notes/%E7%AE%97%E6%B3%95%20-%20%E5%B9%B6%E6%9F%A5%E9%9B%86.md)

# 基础概念
- [数据结构之图的基本概念](https://www.cnblogs.com/xiaobingqianrui/p/8902111.html)

## 存储结构
根据图的**稠密程度**选择存储结构
- 临接矩阵：`[][]bool{}`适合于稠密图，NxN，横纵坐标为对应节点编号
- 临接表：`[]*Node{}`与`map[Node][]*Node{}`适合于稀疏图

## 基本操作
- [Go Data Structures: Graph](https://flaviocopes.com/golang-data-structure-graph/)
- 添加节点：`AddNode()`
- 连接节点形成边：`AddEdge()`

## 搜索
通常需要借助一个标记数组来确认节点是否已经访问过
- BFS：队列
- DFS：栈

# 题型涉及
## 二分图
如果可以用两种颜色对图中的节点进行着色，并且保证相邻的节点颜色不同，那么这个图就是二分图。
## 拓扑排序
常用于在具有先序关系的任务规划中。
## 并查集
并查集可以动态地连通两个点，并且可以非常快速地判断两个点是否连通。
