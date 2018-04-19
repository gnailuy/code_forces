# http://poj.org/problem?id=3181

def solve(N, K):
    f = [0 for _ in range(N+1)]
    f[0] = 1

    for i in range(1, K+1):
        c = 1
        fc = [0 for _ in range(N+1)]
        while True:
            w = c*i
            if w > N:
                break

            for j in range(N, -1, -1):
                if w <= j:
                    fc[j] += f[j-w]
                else:
                    break

            c += 1

        for k in range(1, N+1):
            f[k] += fc[k]

    return f[-1]


with open('poj_3181.data', 'r') as f:
    N, K = [int(x) for x in f.readline().strip().split(' ')]

    print(solve(N, K))

