# http://poj.org/problem?id=3614

class PQMin(object):  # Min 优先队列
    size = 0
    store = None

    def __init__(self, N=10):
        self.store = [None for _ in range(N)]

    def swim(self, i):
        if i > 1:
            parent = i//2
            if self.store[i] < self.store[parent]:
                self.store[i], self.store[parent] = self.store[parent], self.store[i]
                self.swim(parent)

    def sink(self, i):
        a, b = i*2, i*2 + 1
        c = None
        if self.size >= b:
            if self.store[a] < self.store[b]:
                c = a
            else:
                c = b
        elif self.size >= a:
            c = a

        if c is not None and self.store[c] < self.store[i]:
            self.store[c], self.store[i] = self.store[i], self.store[c]
            self.sink(c)

    def enqueue(self, item):
        self.size += 1
        if self.size >= len(self.store):
            self.store.extend([None for _ in range(self.size)])
        self.store[self.size] = item
        self.swim(self.size)
        return self

    def min(self):
        if self.size <= 0:
            return None
        res = self.store[1]
        self.store[1] = self.store[self.size]
        self.size -= 1
        self.sink(1)
        return res

    def peek_min(self):
        if self.size <= 0:
            return None
        return self.store[1]


with open('poj_3614.data', 'r') as f:
    C, L = [int(x) for x in f.readline().strip().split(' ')]
    cl = []
    ll = []
    for _ in range(C):
        cl.append([int(x) for x in f.readline().strip().split(' ')])
    for _ in range(L):
        ll.append([int(x) for x in f.readline().strip().split(' ')])
    cl.sort()  # 按照最小防晒值排序
    ll.sort()  # 按照防晒值排序

    count = 0
    pq = PQMin()
    i = 0
    for l in ll:  # 对于每一瓶防晒霜
        spf = l[0]  # 其防晒值
        cover = l[1]  # 其能涂抹的次数
        while True:
            if i < len(cl) and cl[i][0] <= spf:
                # 最小防晒值大于当前值的，其最大防晒值加入队列
                pq.enqueue(cl[i][1])
                i += 1
            else:
                break

        while True:
            if cover <= 0:
                break
            # 每次取出最小的最大防晒值
            r = pq.min()
            if r is None:
                break

            # 如果最大值大于当前值，则涂一个
            if r >= spf:
                count += 1
                cover -= 1
            # 否则最大值都小于当前值了，后面也就没机会被涂了

    print(count)

