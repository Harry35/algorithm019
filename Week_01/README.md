学习笔记
第一课 数据结构和算法总结
    刻意练习
        过遍数，五毒神掌
            5分钟看题
            直接看解法，看多解法，比较优劣
            背诵、默写好的解法
            今天做，第二天、过一周、面试前一周
        打字练习typing.io
    数据结构
        一维
            基础：字符串，数组，链表
            高级：栈，队列，双端队列，集合，映射，
        二维
            基础：树，图
            高级：二叉搜索树binary search tree，堆heap，并查集disjoint set，字典树Trie
        特殊
            位运算，布隆过滤器
            LRU Cache
    算法
        if else，switch
        for， while loop
        递归 recursion
        搜索：深度优先搜索DFS，广度优先搜索BFS，A*
        动态规划 Dyamic Programming
        二分查找 Binary Search
        贪心 Greedy
        数学， 几何
第二课 2.1 工欲善其事，必先利其器
    工具
        Google
        Mac： iTerm2+zsh
        VSCode
        Leetcode
        code style: Google, Facebook, airbnb
    指法
        fn+delete 删除右边
        cmd+left/right 行头/尾
        option+left/right 光标按单词切分
        option+delete 删除单词
        option+cmd+right 选中整行
        IDE 自动补全，top tips/使用技巧
    自顶向下的编程方式
        代码像新闻稿
        如何做code reivew https://cloud.tencent.com/developer/article/1514271
第二课 2.2 时间复杂度和空间复杂度
    时间复杂度
        所有时间复杂度 https://www.zhihu.com/question/21387264
            O(1) 常数复杂度
            O(logn)对数复杂度
            O (n)线性复杂度
            O(n^2)平方
            O(n^3)立方
            O(2^n)指数
            O(n!)阶乘
        斐波那契数列求n项时间复杂度O(2^n)
        二叉树前中后序遍历时间复杂度O(n)
        二分查找时间复杂度O(logn)
        所有递归函数计算时间复杂度：主定理https://zh.wikipedia.org/wiki/%E4%B8%BB%E5%AE%9A%E7%90%86
    空间复杂度
        数组的长度
        递归的深度
第三课 数组、链表、跳表的基本实现和特性
    数组Array
        时间复杂度
            prepend O(1)
            append O(1)
            insert O(n)
            delete O(n)
            lookup O(1)
    链表Linked list
        时间复杂度
            prepend O(1)
            append O(1)
            insert O(1)
            delete O(1)
            lookup O(n)
    跳表Skip list 
        升维思想+空间换时间
        介绍https://redisbook.readthedocs.io/en/latest/internal-datastruct/skiplist.html
        用来替换平衡树，用于Redis，LevelDB
            为什么Redis使用跳表不是红黑树https://www.zhihu.com/question/20202931
        LRU Cache 
            LRU 缓存机制https://leetcode-cn.com/problems/lru-cache/
        时间复杂度
            prepend O(1)
            append O(1)
            insert O(1)
            delete O(1)
            lookup O(logn)
    实战题
        移动零（华为、字节跳动在近半年内面试常考）
        盛最多水的容器（腾讯、百度、字节跳动在近半年内面试常考）
        爬楼梯（阿里巴巴、腾讯、字节跳动在半年内面试常考）
        三数之和（国内、国际大厂历年面试高频老题）
    作业
        简单
            用 add first 或 add last 这套新的 API 改写 Deque 的代码
            分析 Queue 和 Priority Queue 的源码
            删除排序数组中的重复项（Facebook、字节跳动、微软在半年内面试中考过）
            旋转数组（微软、亚马逊、PayPal 在半年内面试中考过）
            合并两个有序链表（亚马逊、字节跳动在半年内面试常考）
            合并两个有序数组（Facebook 在半年内面试常考）
            两数之和（亚马逊、字节跳动、谷歌、Facebook、苹果、微软在半年内面试中高频常考）
            移动零（Facebook、亚马逊、苹果在半年内面试中考过）
            加一（谷歌、字节跳动、Facebook 在半年内面试中考过）
        中等：
            设计循环双端队列（Facebook 在 1 年内面试中考过）
        困难：
            接雨水（亚马逊、字节跳动、高盛集团、Facebook 在半年内面试常考）
第四课 栈和队列的实现与特性
    栈stack 先入后出
        添加O(1)
        删除O(1)
        查询O(n)
    队列queue 先入先出
        添加O(1)
        删除O(1)
        查询O(n)
    双端队列double end queue，两端可以进出的queue
        添加O(1)
        删除O(1)
    优先队列priority queue
        底层具体实现的数据结构较为多样和复杂：heap、bst、treap
        插入O(1)
        取出O(logn)
    分析Stack、Queue、Priority Queue源码实现
    最近相关性，先来后到的问题用栈解决
    实战
        有效的括号（亚马逊、JPMorgan 在半年内面试常考）
        最小栈（亚马逊在半年内面试常考）
        柱状图中最大的矩形（亚马逊、微软、字节跳动在半年内面试中考过）
        滑动窗口最大值（亚马逊在半年内面试常考）
