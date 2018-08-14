# Algorithms 4e

## GCD

Euclid's Algorithm: 大数除小数；小数变大数，余数为新的小数；循环至小数变零，大数即为 GCD

## SQRT

Newton's Method: 循环 k = (k + x/k)/2，直至 k 和 x/k 相差小于阈值，k 即为 sqrt(x) 的近似值

## Arithmetic Expression (((2 - 5) + 1) * 2)

* Dijkstra's Double Stack: 一个操作符栈，一个操作数栈，可转后缀表达式
* 如果不一定有括号，而且加上乘除，需要考虑操作符优先级，可以通过中缀表达式转逆波兰式来解决
* 用一个栈来保存操作符，扫描中缀表达式，遇到数字则输出，遇到优先级不高于栈顶的符号，或者遇到右括号，则符号出栈加入表达式，否则符号入栈
* 逆波兰式求值：用一个栈从左到右扫，遇到数字入栈，遇到符号出栈计算然后结果入栈

## Editor Buffer

支持插入删除左移右移, 用两个栈头对头，左右移动时一个栈出一个栈进

## Reverse a Linked List

开一个新链表头，模拟旧链表出栈新链表入栈；可以用递归

## 松散链表

1. 链表节点保存一个数组，而不是单个元素
2. 插入时，若数组已满，则 Split 成两个，各占原来一半元素
3. 删除时，若节点数少于一半，则尝试从相邻节点借一些节点过来，若相邻节点也小于一半元素，则合并

## Intersection of two Linked Lists

如果有交叉，那么交叉点之后两个链表就是同一个链表了，长度肯定一样，所以需要两个指针按照剩余长度相同的步调往后查找

## 第 K 个全排列

Hint: 例如 1234 四个数的排列，1 开头的话，234 总共六个排列法，K 如果是 1~6 就是 1 开头，7~12 就是 2 开头

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

## Grey Code

1. Grey Code 编码的相邻两个数之间只有一位不同
2. G = B XOR (B>>1)
3. 或者 G(n) = B(n+1) + B(n) 从低位到高位
4. 或者 B(n) = B(n+1) - G(n) 从高位到低位

``` text
D       G      B
0  　　 000    000
1  　　 001    001
2   　　011    010
3   　　010    011
4   　　110    100
5   　　111    101
6   　　101    110
7   　　100    111
```

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

1. 从数组开头开始，对于第 i 个元素，都和 i(含) 之后的一个随机元素互换位置
2. 实现起来，从数组最后一个元素 i 开始，和 0~i(含) 之间的随机元素互换位置，比较容易写生成随机数的代码

## 水库抽样 Reservoir Sampling

1. 大量元素随机抽取 K 个，首先把前 K 个元素加入大小为 K 的数组
2. 然后从 K+1 个元素开始循环，对于第 i 个元素，生成一个 1~i 之间的随机数 j
3. 如果随机数 j <= K，则以第 i 个元素取代数组中第 j 个元素

## 8 Puzzle

1. 用 A\* 算法，A\* Heuristic Algorithm
2. 初始状态放入优先队列，从初始状态生成所有邻居状态
3. 按照汉明距离或者汉密尔顿距离，所有状态放入优先队列
4. 每次从队列头生成其他邻居状态，直至最终状态

## Boggle backtracking

1. 遍历一个 N\*N 字符数组，每次可上下左右斜线走一步
2. 递归深度遍历，循环每一步可选的邻居节点，循环内递归调用，递归调用之后回退，继续循环

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

### 二分图

1. 一个图为二分图，则可以将所有顶点集合 V 分割为两个子集 A 和 B，使得所有边都满足一个顶点在 A，一个顶点在 B
2. 判断二分图，可以用染色法，从任意顶点开始搜索路径，交替染两种颜色，如果最后成功遍历全图，则是二分图，否则会遇到染色矛盾

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

## 按词翻转句子 (hello world 变 world hello)

翻转字符串，然后每个单词原地翻转

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
* CAP: Consistency, Availability, Partition Tolerance (CP: Zookeeper; AP: DynamoDB; 网络是不稳定的，网络中断导致分区是不可避免的)
* Paxos
* BASE: Base Available, Soft-state, Eventually Consistence

### DB Storage

* B-Tree
* LSM-Tree: 数据 Append 到磁盘 SSTable，并且写 WAL，然后更新内存 MemTable，MemTable 定期固化，通过 compact 操作合并 WAL

### Sharding

* Consistent Hashing: Key 的哈希范围看作环，均匀部署 Sharding 边界，在节点增删失效时，尽量保持节点距离均匀
* Consistent Hashing: 一个实现：根据 Node 名字的 Hash 决定 Node 位置，Node 按 Hash 存在一个 TreeMap 里面，发请求时，根据 Key 的 Hash 找到 TreeMap.tailMap() 中最近的 Node
* 每个 Shard 多个节点，如果考虑 Consistency，那么对多个 Shard 节点的读写需要经过同一的 Master 节点，或者有类似 ZK 的一致性机制

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

* Netty: Bootstrape, EventLoopGroup, Channel, Handler; Reactor, EventLoop, ZeroCopy, JNI
* RESTful
* gRPC

## Cache System

* Redis
* Memcache
* 写入策略：Write Through (同时写), Write Back (写 Cache 然后刷回 DB), Write Around (写 DB)
* 替换策略：LRU, LFU (Frequency)

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
* Cache Coherence: 缓存一致性模型通常对程序员透明，它指的是多核情况下，每个核的 L1 缓存一致性，Write Through 情况下，每个核的 L1 之间要有通知机制，使得其他核知道缓存失效，Write Back 的情况下也是如此，但还要保证数据写回共享缓存或者主存之后才能读，通过写之前标记内存状态来实现
* Memory Consistency: 内存一致性模型主要是为了保证多核情况下读写内存的一致性，程序员的角度看，可以通过原子操作来，如加锁

## Performance

* https://gist.github.com/jboner/2841832

## Design Pattern

### 创建型模式

是对对象创建过程的各种问题和解决方案的总结

1. Factory, AbstractFactory, 对于构造复杂对象，使用复杂的构造函数不好用，用工厂类来提供不同的创建方法、传入创建参数
2. Singleton, 单例，可以在 getInstance() 里 Lazy 创建实例，需要注意线程安全
3. Builder, 通过创建一个 Builder 类来创建对象，Builder 类提供一系列的 set 方法，修改要创建的对象参数

### 结构型模式

是针对软件设计结构的总结，关注于类、对象继承、组合方式

1. Adapter, 创建一个新的类，提供新的 API，但内部调用老的类的 API，用于新接口访问旧对象
2. Decorator, 例如 InputStream，可以包装成 FileInputStream, BufferedInputStream 等
3. Proxy
4. Bridge
5. Composite

### 行为型模式

是从类或对象之间交互、职责划分等角度总结的模式

1. Strategy, 不同类型的请求，在运行时选择用不同的策略或算法去应对
2. Iterator, 创建一个 Iterator 来 traverse 一个容器
3. Observer, 提供一个共用接口被调用，来观察 Runtime 情况
4. Command
5. Interpreter

# Design Note

## Java

### OOP

Encapsulation: Hiding implementation
Abstraction: Using class/interface instead of implementation to interact
Inheritance: "is a" relationship; reuse code
Polymorphism: one name, many forms; compiler ploymorphism and runtime ploymorhpism

### The language

* Write once, run anywhere; JVM; bytecode; JIT Compiler; Mixed Mode
* Generic Typing
* Thread Stack, Heap, Method Area, Constant Pool, Native Stack
* GC: SerialGC, ParNewGC, CMS, ParrallelGC, G1

### Exception and Error

* They both are extend from Throwable
* Exceptions are meant to be catched by the programmer
* Errors are fatal, they usually means a program crash: OOMError, StackOverflowError
* Checked Exceptions and Unchecked Exceptions: Checked/Not Checked at compile time
* Checked Exception: IOException
* Unchecked Exception: NullPointerException, RuntimeException
* Exceptions are side effects

### Reference Type

* Strong Reference: Object is directly referenced from one or more Threads
* SoftReference: JVM will clear SoftReferenced Objects before throwing an OOMError, a good way to implement Cache
* WeakReference: JVM will clear WeakReferenced Objects by GC
* PhantomReference: PhantomReferenced Objects will be enqueue to a ReferenceQueue, user can schedule follow-up cleanup actions

### String, StringBuffer, StringBuilder

* Java String is immutable
* StringBuffer is thread safe
* StringBuilder is not thread safe, but faster

### HashTable, HashMap, TreeMap

* HashTable 比较老，线程安全，速度慢，和 HashMap 类似
* HashMap 是 Bucket + 链表的模式，Bucket 根据 load factor 来自动扩容和 Rehash，链表在过长时会变成树
* LinkedHashMap 是双向链表加哈希，可以实现按照插入顺序遍历，实现 Cache 很方便
* TreeMap 基于 RedBlackTree 实现，可以按照自然大小顺序遍历
* PriorityQueue 也可以按自然大小顺序遍历，基于 Heap 实现

### Wrapper 类型

* 基础数据类型 short int long double float char boolean byte 八个
* 对应都有 Wrapper 类型，Java 提供了 boxing 和 unboxing 的功能，自动转换
* 鉴于程序中常用的数字范围有些，valueOf() 方法是有缓存的，使得 Wrapper 类型使用更高效

### 反射

* 在运行时动态获得对象类型，可以用于实现动态代理、插件机制等，例如动态加载第三方对象
* 动态代理：运行时为对象创建代理对象，用代理对象来操作原对象的方法，代理对象中可以实现通用操作，如日志、安全、资源释放等

### Java 排序

* Arrays.sort() 在 primitive type 上用的是三路 QuickSort
* 在对象类型上用的是归并排序，结合小数组上的插入排序，保证排序稳定

### Vector, ArrayList, LinkedList

* Vector 是线程安全的，底层是数组，自动扩容策略是翻倍
* ArrayList 类似 Vector，但线程不安全，性能好，自动扩容是翻 50%
* LinkedList 也线程不安全，底层是双向链表

## Realtime Chat

1. 客户端使用 Ajax、Comet 等轮询，或者使用 WebSocket 实现双工通信，Server-Sent Events 可以不考虑
2. 服务器端实现一个用户一个 Channel 队列，Channel 中保存用户接受到的消息
3. 用户各客户端已读位置保存在服务器端，在线时，拉取新的消息
4. 用户发送的消息保存在接受方的 Channel 中，由客户端负责解析
5. API 主要包括发消息、获取对话列表、获取对话消息
6. 可以使用对话ID+时间戳的方式保存对话消息，存放在 KV Store 里面，如 HBase

## Twitter 类似服务

1. 客户端主要关注三个写 API：发推、关注、喜欢；一个读 API：获取 Feed 流。其他 API 类似；
2. 涉及到用户表、推文表、关注关系表和喜欢表；
3. Sharding 方案要考虑的比较多，数据库按照用户分区，这样获取一个用户的 Feed 就涉及到跨分区读和合并；
4. 可以在按用户分 Sharding 的基础上，加一层缓存 Layer，缓存用一个大的环形队列+SymbolTable 来做，可以从中获取最新推文；

## Amazon Large Scale

1. EC2: node with web stack
2. Route 53: DNS Service
3. RDS: Relational Database Service
4. DynamoDB: NoSQL KV Service
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

## OAuth 2.0 work flow (rfc6749)

``` text
+--------+                               +---------------+
|        |--(A)- Authorization Request ->|   Resource    |
|        |                               |     Owner     |
|        |<-(B)-- Authorization Grant ---|               |
|        |                               +---------------+
|        |
|        |                               +---------------+
|        |--(C)-- Authorization Grant -->| Authorization |
| Client |                               |     Server    |
|        |<-(D)----- Access Token -------|               |
|        |                               +---------------+
|        |
|        |                               +---------------+
|        |--(E)----- Access Token ------>|    Resource   |
|        |                               |     Server    |
|        |<-(F)--- Protected Resource ---|               |
+--------+                               +---------------+
```

## Java 并发编程 (synchronized and Lock)

1. 重量级锁，通过占用 monitor 来实现，依赖操作系统 Mutex，需要切换内核态
2. CAS(Compare And Swap): 修改之前先判断值是否改变
3. 偏向锁，只有本线程占用对象时，直接执行同步区域，有线程竞争时则撤销偏向，使用轻量级锁尝试锁定
4. 轻量级锁，尝试用 CAS 方法占用对象，成功的情况下先使用轻量级锁，失败时则膨胀为重量级锁
5. 自旋锁，用短时的忙等待期待资源迅速被释放，自旋次数根据自旋成功与否自适应改变
6. 循环 wait()，防止被无效唤醒，修改了数据之后要 notify() 或 notifyAll() 唤醒其他线程
7. volatile: 禁止指令重排优化，保证变量值变化时所有线程都立即可见
8. 实现 singleton 的一个方案，利用静态子类来持有 singleton 对象，getInstance() 返回子类的静态 instance 对象，ClassLoader 保证了对象唯一性，又是线程安全的，避免了同步代码

## Java 并发库

* 老的同步容器，读写全部同步: HashTable, Vector, Stack
* 同步 Wrapper，对象级锁: SynchronizedMap: Map\<String, String> synchronizedHashMap = Collections.synchronizedMap(new HashMap\<String, String>());
* 并发容器，提供分 Segment 的锁，Java 8 升级为 Bucket 锁: ConcurrentHashMap, CopyOnWriteArrayList
* 线程安全队列: ArrayBlockingQueue, SynchronousQueue

## IO 和并发模型

1. 传统同步 IO，InputStream/OutputStream 抽象了字节流，Reader/Writer 则加上了字符编码解码能力，Buffered Stream 增加了缓冲区
2. Java NIO 中 Selector 监听一堆 Channel，例如 SocketChannel，Channel 把数据读到 Buffer，Channel Ready 后交给 Thread 来处理
3. NIO2 支持 Callback 方式的异步 IO，为 Channel 注册 Handler，IO 事件完成后回调 Handler
4. NIO 提供的 Channel 之间 transferTo() 的方法，利用了操作系统 Zero Copy 技术，无需把内核缓冲区数据拷贝回用户缓冲区，比 I/OStream 对拷效率高
5. DirectByteBuffer 利用堆外内存，可以保存生命周期较长的对象，可以用于 Zero Copy IO，可以规避 GC 暂停，这些内存通过 PhantomReference 和 Cleaner 机制回收

* Delegator + Workers 模式：Worker 以 Blocking 的方式运行，需等待 IO，Worker 线程遇到慢 IO 时可能被耗光
* Reactor + Channel 模式：Worker 以 Non-blocking 模式运行，遇到 IO 注册 IO 完成事件，然后交出控制权回到线程池

## 负载均衡和反向代理

1. DNS Round Robin (rarely used)
2. L3/L4 Layer (TCP/IP)
3. L7 Layer (Application layer, e.g. HTTP)
4. 服务端通过 Reverse Proxy 将后端 API 整合到统一的出口

## TCP handshake

``` text
    TCP A                                                 TCP B
1.  CLOSED                                                LISTEN
2.  SYN-SENT    --> <SEQ=100><CTL=SYN>                --> SYN-RECEIVED
3.  ESTABLISHED <-- <SEQ=300><ACK=101><CTL=SYN,ACK>   <-- SYN-RECEIVED
4.  ESTABLISHED --> <SEQ=101><ACK=301><CTL=ACK>       --> ESTABLISHED
5.  ESTABLISHED --> <SEQ=101><ACK=301><CTL=ACK><DATA> --> ESTABLISHED
```

# Hadoop and Spark Review

1. reduce() 接受一个函数，参数是两个 RDD 中的元素，返回和参数同类型的聚合值
2. fold() 多接受一个零值，函数将 RDD 里的值和零值合并
3. aggregate() 又多接受一个函数，一个函数负责零值和 RDD 元素的合并，另一个函数负责两个 RDD 元素的合并，因此零值类型可以和元素类型不同
4. map() 是个 transformation，它的函数将每个元素进行一个转换并返回新元素，而 foreach() 是个 action，它不返回新 RDD，函数主要是 Side effect
5. countByValue() 类似于 bash 的 uniq -c
6. 主要集合操作 union(), intercection(), substract(), cartesian()
7. reduceByKey() 用于将 Pair RDD 按照 Key 进行 reduce；类似的 groupByKey() 按 Key 聚合成集合；foldByKey() 按 Key 来 fold()
8. combineByKey() 和 aggregate 类似，按照 Key 进行元素合并，首先在分区内提供将元素映射成累加值，以及累加值和元素合并两个方法，分区间提供累加值互相合并的方法
9. mapValues() 相当于 map() 只修改 Value；flatMapValues() 和 flatMap() 类似，都将传入的 Value 转换成 Iterable
10. 主要 KV RDD 相交操作 join(), left(right)OuterJoin(), cogroup()
11. cogroup() 将两个 Pair RDD 按照 Key 合并，每个 Key 对应两个 Iterable，分别是原 RDD 的元素集合，join() 基于它实现
12. mapPartitions(WithIndex)() 和 foreachPartitions() 分别都在分区内进行循环
13. pipe() 用来和外部脚本、程序进行对接

