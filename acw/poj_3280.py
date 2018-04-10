# http://poj.org/problem?id=3280

def mirror_cost(s, v, c_cost, i, j):
    if i < 0:
        return sum([c_cost[x] for x in s[j:]])
    elif j >= len(s):
        return sum([c_cost[x] for x in s[:i+1]])
    else:
        if v[i][j] is None:
            if s[i] == s[j]:
                v[i][j] = mirror_cost(s, v, c_cost, i-1, j+1)
            else:
                v[i][j] = min(mirror_cost(s, v, c_cost, i-1, j) + c_cost[s[i]],
                              mirror_cost(s, v, c_cost, i, j+1) + c_cost[s[j]])
        return v[i][j]


with open('poj_3280.data', 'r') as f:
    n, m = [int(x) for x in f.readline().strip().split(' ')]
    s = f.readline().strip()
    c_cost = {}
    for _ in range(n):
        xc = f.readline().strip().split(' ')
        c_cost[xc[0]] = min([int(x) for x in xc[1:]])

    v = [[None for _ in range(m)] for _ in range(m)]

    m_cost = None
    for i in range(m+1):
        l_cost = min(mirror_cost(s, v, c_cost, i-1, i),
                     mirror_cost(s, v, c_cost, i-1, i+1))
        if m_cost is None or m_cost > l_cost:
            m_cost = l_cost

    print(m_cost)

