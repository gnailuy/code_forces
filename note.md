# Algorithms 4e

## GCD

Euclid's Algorithm: 循环大数除小数

## SQRT

Newton's Method: 循环 xNew = ( xOld + x/xOld )/2，直至 xOld 和 x/xOld 相差小于阈值

## Arithmetic Expression (((2 - 5) + 1) * 2)

Dijkstra's Double Stack: 一个操作符栈，一个操作数栈，可转后缀表达式

## Editor Buffer

支持插入删除左移右移, 用两个栈头对头，左右移动时一个栈出一个栈进

## Reverse a Linked List

开一个新链表头，模拟旧链表出栈新链表入栈；可以用递归

## 调和级数

1 + 1/2 + 1/3 + ... + 1/N ~ lnN

## 斯特灵估计

lgN! ~ NlgN: 该估计从斯特灵公式得来

## 二项式系数

当 k 相对小时 C(N, k) ~ N^k/k!，其中 N(N-1)(N-2)...(N-k+1) ~ N^k

## (1-1/x)^x ~ 1/e

## 一维数组 N 个不同的数字，找到一个局部最小值

使用二分搜索，因为若中点非局部最小，则小的一侧至少存在一个局部最小

## 测试鸡蛋从几层楼摔下不会破的问题

1. 如果鸡蛋很多，可以二分搜索
2. 如果只有两个鸡蛋，可以按照 sqrt(N), 2sqrt(N), ..., sqrt(N)sqrt(N) 测试第一个，测到碎了之后另一个线性扔

## 选择排序和插入排序

1. 选择排序运行时间和输入顺序无关
2. 插入排序在输入基本有序的情况下更优
3. 输入基本有序时，优先考虑插入排序 (下面 Shell 排序的优化依据)

## Shell 排序

1. 插入排序的一个扩展，每次把相隔 h 的元素组成的子序列插入排序，大循环内减少 h，当 h 为 1 时全局就已排序
2. h 可以选择小于 N/3 的数，按照从 1 开始，h=3h+1 递增找到初始 H，大循环内按照 h=h/3 递减 h

## MergeSort

1. 需要一块额外数组，轮换角色法可以减少数组来回拷贝次数
2. 另一个优化是子数组长度小时，使用插入排序，一般长度小于 15 时使用插入排序
3. Merge 之前检查要 Merge 两个数组，如果一个的最大值小于另一个的最小值，则可以省去一次 Merge
4. MergeSort 复杂度可以用比较的次数来衡量，每次循环比较次数上限均为 N，循环次数为 lgN，因此复杂度 NlgN
5. 基于比较的排序比较次数不能低于 NlgN
6. 自底向上写循环进行归并排序，每 2、4、8... 分组依次向上归并
7. 可以处理链表的排序，使用一个队列，初始队列中是 N 个单元素队列，循环每次取前两个队列合并，合并后加入队尾

## QuickSort

1. 和 MergeSort 相比：前者先按照参考值把数组划分成两部分，然后递归排序两部分；后者先对半划分，然后递归排序两部分，最后合并两部分
2. QuickSort 的最优情况是每次参考值恰好把数组对半分
3. 优化：长度小的数组用插入排序，或忽略小数组，最后统一插入排序一次
4. 优化：每次选取一个小样本，用样本中值作为参考值
5. 优化：对于存在大量重复 Key 的情况，可以三路分区，选取参考值后，将数组分为小于、等于、大于三部分；3-Way QuickSort 通常是排序库函数实现
6. 用栈来实现非递归的 QuickSort，每次 Devide 之后，把大小两个子数组压栈，中间压入参考值

### 一个基于 QuickSort 方法的选择第 K 个元素算法

1. 先 shuffle 数组
2. 然后每次用 QuickSort 同样的办法进行 partition
3. 如果 partition 得到的分界 j 即为 K，则返回元素
4. 否则如果 j 小于 K，在 j+1 及之后的数组上重复算法
5. 如果 j 大于 K，在开头到 j-1 的数组上重复算法

## 稳定性

1. 插入排序和归并排序是稳定的，选择、希尔、快速、堆排序是不稳定的
2. 比较元素时先比较值，再比较其 index，可以把不稳定排序变成稳定排序

## PQ and Heap Sort

1. 用数组表示的完全二叉树实现大顶、小顶堆，从下标 1 开始存数据编码方便很多，实现优先队列 PQ 很方便
2. Index PQ 实现 N 路 stream 归并：用大小为 N 的 PQ，每一路 stream 存一个元素进来，每次输出堆顶元素之后，再从对应 stream 里读新元素进来
3. 堆排序过程：自右向左，从 N/2 到 1 为堆顶，用 sink() 构造大顶堆；然后每次取出堆顶元素放在堆末尾，堆大小减一，原来末尾的元素放在堆顶进行 sink()

## Shuffle 的一个简单实现

从数组开头开始，对于第 i 个元素，都和 i(含) 之后的一个随机元素互换位置

## 8 Puzzle

1. 用 A\* 算法，A\* Heuristic Algorithm
2. 初始状态放入优先队列，从初始状态生成所有邻居状态
3. 按照汉明距离或者汉密尔顿距离，所有状态放入优先队列
4. 每次从队列头生成其他邻居状态，直至最终状态

## Binary Search Tree

### BST 的删除

1. 找到要删除的节点 x，如果子树少于两棵，直接用子树替代 x 即可
2. 如果有两棵子树，提升左子树中最大者，或者右子树中最小者

## Red-Black Tree BST

1. Red-Black Tree BST 的插入：插入的节点为红色，红色向上传递，用递归实现（左黑右红左旋；左红红右旋；左右红反转），最后把根节点标记为黑色
2. Red-Black Tree BST 删除最小：自上而下搜索时，调整当前节点不是黑色，删除最小之后，自下而上重新调整不合规范的红色节点
3. Red-Black Tree BST 的删除：基于删除最小实现，如果要删除的节点是叶节点，直接子树删除最小，否则用右子树最小值替代本节点，删除右子树最小
4. Red-Black Tree BST 的插入、删除、查找等操作均可达到对数复杂度，是常用的 Symbol Table 实现

## Hash

1. 取模做 Hash 时，模要取素数，否则被模的元素中总有部分数位无法被利用
2. Hash 冲突常见解决办法一：每个 Hash 位都存放一个链表ST，叫做 Separate Chaining
3. Hash 冲突常见解决办法二：冲突之后向后寻找空位，此时需要维护数组长度，经常 Resize

## 用 SymbolTable 实现 List

1. 一个 ST 用来查找索引 i 到 item
2. 另一个 ST 用来查找 item 本身

## UniQueue

实现一个 Queue，并用一个 ST 来追踪 Queue 中的元素，enqueue 和 dequeue 的时候分别查询和删除 ST

## Union-Find

1. 先构建图，然后用搜索方法在图中找出联通区域
2. 用第一章的加权森林法，用数组表示链接，c[i]=j 表示 i 指向 j，再用数组保存距离跟节点的距离，总是把小树连到大树上

## 有向图

### 有向图环检测

1. 用深度优先遍历，遍历时维护一个 marked 数组和一个 onStack 数组
2. 把遇到的点都标记为 marked 和 onStack，在处理完一个点之后取消 onStack 标记(递归结束时)
3. 遍历过程中遇到已经 onStack 的点，则说明有环
4. 如果要返回具体环的路径，则记录一个 edgeTo 数组，追踪当前点是从那里指向来的，检测到环时按照 edgeTo 恢复路径

### 有向图拓扑排序

1. 深度优先遍历，从任意节点开始遍历都可以
2. 在递归调用之后将节点 push 进一个 stack(先递归调用，后 push 节点)
3. 最后弹出 stack 即为拓扑排序的一个结果

### 有向图强连通检测

1. 首先，在 Reverse 图上找一个拓扑排序顺序；然后，按照该顺序进行深度优先遍历，遍历可达点均为联通区域
2. Reverse 图上面的拓扑排序中，若 v 排在 w 后面，则说明在原图中 v 到 w 可达，而如果从 w 开始深度遍历访问到了 v，则 v 和 w 就一定位于同一个强连通区域里

### 加权有向图的最短路径 Dijkstra 算法(非负权重)

1. 从源点 s 开始，distTo[] 数组表示各点到 s 的距离，初始 distTo[s] 为零，其它为正无穷
2. 将可达的边按照 distTo[] 排序加入优先队列，每次取最短可达的边，relax 这个边，然后将该边 to 顶点可达的边也加入优先队列
3. relax 一个边的操作：对于一个边的 from 点 v 和 to 点 w，如果 distTo[v] + edge < distTo[w]，那么 更新 distTo[w]

## 无向图

### 加权无向图的最小生成树贪心规则

对于图的任意一个子图，连接子图和子图的补图的所有边中，权重最小的边一定在生成树里

### 加权无向图的最小生成树 Prim 算法

1. 从任意节点开始构造子图，用优先队列保存该子图所有对外的边
2. 选择最小的边，把对应的节点加入子图，可达的边加入优先队列，直至所有节点都已经进来

### 加权无向图的最小生成树 Kruskal 算法

1. 用优先队列保存所有边，每次选择最小的边，如果引入该边不构成环，则引入
2. 直至选够节点数 V-1 个边位置
3. 其中构成环判断可以用 Union-Find 算法

## Key-Indexed Counting

适用于 Sorting Key 为小范围整数的情况

1. 遍历对 Key 计数生成 cnt[]
2. 通过 cnt[i] += cnt[i-1] 计算每个 Key 的起始点
3. 遍历，根据起始点放置元素
4. 按原有元素顺序拷贝回原数组，实现排序

## 等长字符串排序

1. 字符串长度为 K，则从右到左做 K 次 Key-Indexed Counting 排序
2. 因为 Key-Indexed Counting 排序是稳定的，所以可行

## 任意字符串排序

1. 从前向后按照字符进行 Key-Indexed Counting 排序
2. 然后对于相同开头的字符串，递归排序其剩余部分
3. 对于小的子字符串集合，切换为插入排序来题高性能

## Trie

1. 实现一种字符串为 Key 的 Symbol Table，使用树形结构 Share Key Prefix
2. 每个节点上的 N(字符集) 维数组指向后继节点，节点有值表示对应的 Key 有 Value

## Ternary Search Tries

1. 每个节点有一个值，表示当前字符
2. 有三个链接，左边指向比当前字符小的节点，右边指向比当前字符大的节点，中间则指向当前字符顺下去的节点

## KMP 字符串搜索算法

1. 为 Pattern 字符串构建一个 dfa
2. dfa[c][j] 代表遇到字符 c 时，如果位于 Pattern 第 j 个字符，下一步应该对比 Text(i+1) 和 Pattern(dfa[c][j])
3. 以达到 M+N 的最坏搜索复杂度

## 压缩算法

### 字符集有限

如 DNA 数据，用更小的字符集来表示数据

### Run-length 压缩

长的 0 和 1 串较多时，如图像，可以用 0 和 1 的计数值来压缩

### Huffman compression

1. 第一遍输入数据求各字符频率，按照频率从小到大排序，当作单节点 Trie，放入优先队列
2. 每次取最小频率的两个子 Trie:，频率相加合并为新的 Trie 加入优先队列

## Event-driven simulation

1. 用优先队列保存事件，按照时间顺序依次发生
2. 事件的发生可能导致其他事件变为无效，可以惰性判断，也可以通过关联 SymbolTable 主动标记

## B-Tree

1. 每个节点有 N 个元素，每个元素为一个指针，指向不大于该元素的子节点
2. 第一个元素可以用 * 号表示占位符，所有元素都不大于占位符

## Suffix Array

1. 将字符串的所有 Suffix 求出，组成一个 Array
2. 排序这个 Suffix Array，即可快速求出最大重复字串

## Network Flow

任意一个 st-cut，也就是分开源点 s 和终点 t 的 cut，穿越 cut 的所有边 flow 之和即为网络的 flow；

# Skill Tree

## Toolkit

* Python
* Java/Scala
* Golang
* C
* Scripting
* Git

## ETL

* Yarn
* MapReduce
* Spark
* Spark Streaming
* HBase
* HDFS

## Database

### Relational DB

* MySQL
* PostgreSQL
* SQL
* Partition
* Replication

### NoSQL DB

* Big Table(Column-based): HBase
* Key Value DB: Riak
* Document DB: MongoDB
* Search: Elastic
* Graph DB: Neo4j
* ACID: Atomicity, Consistency, Isolation, Durability
* CAP: Consistency, Availability, Partition Tolerance
* BASE: Base Available, Soft-state, Eventually Consistence

## Queue

* Kafka
* NSQ
* Redis Queue

## Web Development

* MVC
* Spring Boot
* Flask/Django
* Beego

## Frontend

* HTML/CSS
* Javascript

## HTTP and RPC

* RESTful
* gRPC

## Cache System

* Redis
* Memcache

## Moniting and Logging

* Ganglia/Zabbix
* Log Rotation
* Logstash
* WAL (Write Ahead Log)

## Concurrency

* Non-blocking
* Asynchronous
* Locking
* Resource Decouple

## Performance

* https://gist.github.com/jboner/2841832

# Design Note

## Realtime Chat

1. 客户端使用 Ajax、Comet 等轮询，或者使用 WebSocket 实现双工通信
2. 服务器端实现一个用户一个 Channel 队列，Channel 中保存用户接受到的消息
3. 用户各客户端已读位置保存在服务器端，在线时，拉取新的消息
4. 用户发送的消息保存在接受方的 Channel 中，由客户端负责解析

## Amazon Large Scale

1. EC2: node with web stack
2. Route 53: DNS Service
3. RDS: Relational Database Service
4. DynamoDB: NoSQL Service
5. Redshift: Data Warehouse
6. ELB: Load Balance
7. CloudFront: CDN
8. S3: Object Store (like Aliyun OSS)
9. ElastiCache: Redis or Memcache Service
10. SQS: Queue Service

## DB Large Scale

1. Federation: Split functions into different DB
2. Sharding: Split DB onto multiple hosts
3. NoSQL

## Session

1. Use session-aware load balancer
2. Store session data in distributed cache, like Redis or Memcache

## JIT

1. Java Runtime 使用 JIT 将频繁被执行的字节码编译为执行更快的机器码，从而优化运行速度
2. Python 的 PyPy 也有一个 JIT，所以比 CPython 快很多

