# CREATED
2019-07-15

# 目录说明
目录内容涉及【单机/网络】并发外部排序算法实现。
```sh
.
├── external            # 外部排序算法wrapper
│   └── sort.go
├── main.go             # 入口文件
└── pipeline            # 算法pipeline实现
    ├── naive.go        # 单机版本数据读/写，及算法pipeline实现
    ├── naive_test.go   # 对应测试代码
    ├── network.go      # 网络版本数据读/写
    └── util.go         # 一些工具函数代码
```

# 学习资料
- 面试中的大数据排序问题
    - [大数据算法：对5亿数据进行排序](https://blog.csdn.net/lemon_tree12138/article/details/48783535)介绍了分治法与字典树算法
    - [海量数据排序——如果有1TB的数据需要排序，但只有32GB的内存如何排序处理](https://blog.csdn.net/FX677588/article/details/72471357)
- 博客/出版物
    - ccmouse：[搭建并行处理管道，感受GO语言魅力](https://www.imooc.com/learn/927)
    - [Parallel Merge Sort in Go](https://hackernoon.com/parallel-merge-sort-in-go-fe14c1bc006)
- 其它代码实现
    - [oxtoacart/emsort](https://github.com/oxtoacart/emsort)
    - [azrle/go-external-sort](https://github.com/azrle/go-external-sort)
    - [siddhant2408/external-merge-sort](https://github.com/siddhant2408/external-merge-sort)
