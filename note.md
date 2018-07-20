# Algorithms 4e

1. GCD: Euclid's Algorithm, 循环大数除小数
2. SQRT: Newton's Method, 循环 xNew = ( xOld + x/xOld )/2，直至 xOld 和 x/xOld 相差小于阈值
3. Arithmetic Expression (((2 - 5) + 1) * 2): Dijkstra's Double Stack, 一个操作符栈，一个操作数栈，可转后缀表达式
4. Editor Buffer: 支持插入删除左移右移, 用两个栈头对头，左右移动时一个栈出一个栈进
5. Reverse a Linked List: 开一个新链表头，模拟旧链表出栈新链表入栈；可以用递归
6. 调和级数 1 + 1/2 + 1/3 + ... + 1/N ~ lnN
7. 斯特灵估计 lgN! ~ NlgN，该估计从斯特灵公式得来
8. 二项式系数，当 k 相对小时 C(N, k) ~ N^k/k!，其中 N(N-1)(N-2)...(N-k+1) ~ N^k
9. (1-1/x)^x ~ 1/e
10. 一维数组 N 个不同的数字，找到一个局部最小值：使用二分搜索，因为若中点非局部最小，则小的一侧至少存在一个局部最小；
11. 测试鸡蛋从几层楼摔下不会破的问题，如果鸡蛋很多，可以二分搜索；如果只有两个鸡蛋，可以按照 sqrt(N), 2sqrt(N), ..., sqrt(N)sqrt(N) 测试第一个，测到底线之后另一个线性扔；
12. 选择排序和插入排序，前者运行时间和输入顺序无关，后者在输入基本有序的情况下更优；输入基本有序时，优先考虑插入排序；
13. Shell 排序是插入排序的一个扩展，每次把相隔 h 的元素组成的子序列排序，大循环内减少 h，当 h 为 1 时全局就已排序；h 可以选择 h=3h+1 这样的序列递减；
14. MergeSort 需要一块额外数组，轮换角色法可以减少数组来回拷贝次数；另一个优化是子数组长度小时，使用插入排序，一般长度小于 15 时使用插入排序；Merge 之前检查要 Merge 两个数组，如果一个的最大值小于另一个的最小值，则可以省去一次 Merge；
15. MergeSort 复杂度可以用比较的次数来衡量，每次循环比较次数上限均为 N，循环次数为 lgN，因此复杂度 NlgN；此外，基于比较的排序比较次数不能低于 NlgN；
16. 自底向上写循环进行归并排序，每 2、4、8... 分组依次向上归并，还可以处理链表的排序；还可以使用一个队列，初始队列中是 N 个队列，各含一个元素，循环每次取前两个队列合并；
17. QuickSort 和 MergeSort 相比：前者先按照参考值把数组划分成两部分，然后递归排序两部分；后者先对半划分，然后递归排序两部分，最后合并两部分；
18. QuickSort 的最优情况是每次参考值恰好把数组对半分；几个优化：长度小的数组用插入排序，或忽略小数组，最后统一插入排序一次；每次选取一个小样本，用样本中值作为参考值；
19. 对于存在大量重复 Key 的情况，QuickSort 的一个优化是进行三路分区，选取参考值后，将数组分为小于、等于、大于三部分；3-Way QuickSort 通常是排序库函数实现；
20. 用栈来实现非递归的 QuickSort，每次 Devide 之后，把大小两个子数组压栈，中间压入参考值；
21. 用数组表示的完全二叉树实现大顶、小顶堆，从下标 1 开始存数据编码方便很多，实现优先队列 PQ 很方便；
22. Index PQ 实现 N 路 stream 归并：用大小为 N 的 PQ，每一路 stream 存一个元素进来，每次输出堆顶元素之后，再从对应 stream 里读新元素进来；
23. 堆排序过程：自右向左，从 N/2 到 1 为堆顶，用 sink() 构造大顶堆；然后每次取出堆顶元素放在堆末尾，堆大小减一，原来末尾的元素放在堆顶进行 sink()；
24. 插入排序和归并排序是稳定的，选择、希尔、快速、堆排序是不稳定的；比较元素时先比较值，再比较其 index，可以把不稳定排序变成稳定排序；
25. shuffle 的一个简单实现：从数组开头开始，对于第 i 个元素，都和 i 之后的一个随机元素互换位置；
26. 一个基于 QuickSort 方法的选择第 K 个元素算法：先 shuffle 数组，然后每次用 QuickSort 同样的办法进行 partition，如果 partition 得到的分界 j 即为 K，则返回元素，否则如果 j 小于 K，在 j+1 及之后的数组上重复，如果 j 大于 K，在开头到 j-1 的数组上重复算法；
27. 8 puzzle: 用 A\* 算法，初始状态放入优先队列，从初始状态生成所有邻居状态，按照汉明距离或者汉密尔顿距离放入优先队列，每次从队列头生成邻居状态，直至最终状态；
28. Binary Search Tree BST 的删除：找到要删除的节点 x，如果子树少于两棵，直接用子树替代 x 即可；如果有两棵子树，提升左子树中最大者，或者右子树中最小者；
29. Red-Black Tree BST 的插入：插入的节点为红色，红色向上传递，用递归实现（左黑右红左旋；左红红右旋；左右红反转），最后把根节点标记为黑色；
30. Red-Black Tree BST 的删除：基于删除最小实现，如果要删除的节点是叶节点，直接子树删除最小，否则用右子树最小值替代本节点，删除右子树最小；
31. Red-Black Tree BST 删除最小：自上而下搜索时，调整当前节点不是黑色，删除最小之后，自下而上重新调整不合规范的红色节点；
32. Red-Black Tree BST 的插入、删除、查找等操作均可达到对数复杂度，是常用的 Symbol Table 实现；
33. 取模做 Hash 时，模要取素数，否则被模的元素中总有部分数位无法被利用；
34. Hash 冲突两个常见解决办法：一个是每个 Hash 位都存放一个链表ST，叫做 Separate Chaining；另一个是冲突之后向后寻找空位，此时需要维护数组长度，经常 Resize；
35. 用 SymbolTable 实现 List，一个 ST 用来追踪索引 i 到 item；另一个 ST 用来追踪 item 本身；
36. UniQueue，实现一个 Queue，并用一个 ST 来追踪 Queue 中的元素，enqueue 和 dequeue 的时候分别查询和删除 ST；
37. Union-Find: 先构建图，然后用搜索方法在图中找出联通区域；第一章的加权森林法，用数组表示链接，c[i]=j 表示 i 指向 j，再用数组保存距离跟节点的距离，总是把小树连到大树上；
38. 有向图环检测：用深度优先遍历，遍历时维护一个 marked 数组和一个 onStack 数组，把遇到的点都标记为 marked 和 onStack，在处理完一个点之后取消 onStack 标记；遍历过程中遇到已经 onStack 的点，则说明有环；如果要返回具体环的路径，则记录一个 edgeTo 数组，追踪当前点是从那里指向来的，检测到环时按照 edgeTo 恢复路径；
39. 有向图拓扑排序：深度优先遍历，在递归调用之后将节点 push 进一个 stack，最后弹出 stack 即为拓扑排序的一个结果；从任意节点开始遍历都可以；
40. 有向图强连通检测：首先，在 Reverse 图上找一个拓扑排序顺序，然后，按照该顺序进行深度优先遍历；Reverse 图上面的拓扑排序中，若 v 排在 w 后面，则说明在原图中 v 到 w 可达，而如果从 w 开始深度遍历访问到了 v，则 v 和 w 就一定位于同一个强连通区域里；
41. 加权无向图的最小生成树贪心规则：对于图的任意一个子图，连接子图和子图的补图的所有边中，权重最小的边一定在生成树里；
42. 加权无向图的最小生成树 Prim 算法：从任意节点开始构造子图，用优先队列保存该子图所有对外的边，选择最小的边，把对应的节点加入子图，直至所有节点都已经进来；
43. 加权无向图的最小生成树 Kruskal 算法：用优先队列保存所有边，每次选择最小的边，如果引入该边不构成环，则引入，直至选够节点数 V-1 个边位置；其中构成环判断可以用 Union-Find 算法；
44. 加权有向图的最短路径 Dijkstra 算法(非负权重)：从源点 s 开始，distTo[] 数组表示各点到 s 的距离，初始 distTo[s] 为零，其它为正无穷；将可达的边按照 distTo[] 排序加入优先队列，每次取最短可达的边，relax 其可达的顶点，然后将该顶点可达的边也加入优先队列；
45. Key-Indexed Counting 适用于 Sorting Key 为小范围整数的情况：1. 遍历对 Key 计数生成 cnt[]；2. 通过 cnt[i] += cnt[i-1] 计算每个 Key 的起始点；3. 遍历，根据起始点放置元素；4. 元素拷贝回原数组，实现排序；
46. 等长字符串排序：字符串长度为 K，则从右到左做 K 次 Key-Indexed Counting 排序；因为 Key-Indexed Counting 排序是稳定的，所以可行；
47. 任意字符串排序：从前向后按照字符进行 Key-Indexed Counting 排序，然后对于相同开头的字符串，递归排序其剩余部分；对于小的子字符串集合，切换为插入排序来题高性能；
48. Trie: 实现一种字符串为 Key 的 Symbol Table，使用树形结构 Share Key Prefix；每个节点上的 N 维数组指向后继节点，节点有值表示对应的 Key 有 Value；
49. Ternary Search Tries: 每个节点有一个值，表示当前字符；有三个链接，左边指向比当前字符小的节点，右边指向比当前字符大的节点，中间则指向当前字符顺下去的节点；
50. KMP 字符串搜索算法：首先，为 Pattern 字符串构建一个 dfa，dfa[c][j] 代表遇到字符 c 时，如果位于 Pattern 第 j 个字符，下一步应该对比 Text9(i+1) 和 Pattern(dfa[c][j])；以达到 M+N 的最坏搜索复杂度；

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

