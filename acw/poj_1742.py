# http://poj.org/problem?id=1742

def solve(n, m, a_list, c_map):
    # 滚动数组，f[i] 表示考虑当前物品时，组成金额 i 是否有可能
    # 初始没有物品，只有金额 0 有可能
    f = [False for _ in range(m+1)]
    f[0] = True

    # 遍历所有金额
    for a in a_list:
        fc_list = []
        # 考虑取该金额硬币所有个数的可能性
        for i in range(1, c_map[a]+1):
            w = a*i  # 取 i 个金额为 a 的硬币
            fc = f.copy()
            for v in range(m, 0, -1):
                # 构成金额 v，v 大于等于当前硬币金额 w 时，可能性等于 f[v-w](取用 w) 或 f[v](不取用 w)
                # v 小于 w 时，可能性就是 f[v](不取用 w)，不用更新
                if v >= w:
                    fc[v] = f[v-w] or f[v]
            # 金额 a 的硬币每一种可能的取法保存到一个新的可能性数组里
            fc_list.append(fc)

        # 合并金额 a 硬币的所有可能性数组，更新数组 f
        for fc in fc_list:
            for i in range(m+1):
                f[i] = f[i] or fc[i]

    return sum(f[1:])  # 金额零去掉，其他数组为 True 的为可能金额，求 sum 即可


with open('poj_1742.data', 'r') as f:
    while True:
        n, m = [int(x) for x in f.readline().strip().split(' ')]
        if n == 0 and m == 0:
            break

        ac_list = [int(x) for x in f.readline().strip().split(' ')]
        a_list = ac_list[:n]
        c_map = dict([(a_list[i], ac_list[n+i]) for i in range(n)])

        a_list.sort()
        a_list.reverse()

        print(solve(n, m, a_list, c_map))

