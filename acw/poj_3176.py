# http://poj.org/problem?id=3176

with open('poj_3176.data', 'r') as f:
    n = int(f.readline().strip())

    traingle = []
    for _ in range(n):
        l = [int(x) for x in f.readline().strip().split(' ')]
        traingle.append(l)

    states = [[0 for _ in range(i+1)] for i in range(n)]
    states[0][0] = traingle[0][0]

    for i in range(1, n):
        for j in range(len(traingle[i])):
            if j == 0:
                states[i][j] = states[i-1][j] + traingle[i][j]
            elif j == len(traingle[i]) - 1:
                states[i][j] = states[i-1][j-1] + traingle[i][j]
            else:
                states[i][j] = max(states[i-1][j-1] + traingle[i][j],
                                   states[i-1][j] + traingle[i][j])

    print(max(states[-1]))

