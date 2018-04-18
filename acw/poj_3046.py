# http://poj.org/problem?id=3046

def solve(T, A, S, B, ants):
    f = [0 for _ in range(B+1)]  # 滚动数组，组成大小为 i 的集合的组合数
    for a in range(1, T+1):  # 循环考虑每一个 Family 的蚂蚁
        c = ants[a]  # 蚂蚁个数
        for j in range(B, 0, -1):  # 从后向前，考虑集合大小 j
            for i in range(1, c+1):  # 考虑选取当前蚂蚁为 i 个时
                if j-i > 0 and f[j-i] > 0:
                    f[j] += f[j-i]  # 集合大小 j 的组合数加上集合大小 j-i 的组合数
        for i in range(1, min(c+1, B+1)):  # 考虑只放本组蚂蚁，没有别的蚂蚁时
            f[i] += 1

    return sum(f[S:])


with open('poj_3046.data', 'r') as f:
    T, A, S, B = [int(x) for x in f.readline().strip().split(' ')]
    ants = [0 for _ in range(T+1)]
    for _ in range(A):
        ants[int(f.readline())] += 1

    print(solve(T, A, S, B, ants))

