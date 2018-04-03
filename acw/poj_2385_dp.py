# http://poj.org/problem?id=2385

def max_apple(schedule, t, w):
    dp = [[0 for _ in range(w+1)] for _ in range(t+1)]

    for i in range(1, t+1):
        s = schedule[i-1]
        dp[i][0] = dp[i-1][0] + (1 if s == 0 else 0)  # Don't move
        for j in range(1, w+1):
            dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]) + (1 if s == j%2 else 0)  # Move j steps

    return max(dp[t])


with open('poj_2385.data', 'r') as f:
    t, w = [int(x) for x in f.readline().strip().split(' ')]
    schedule = []
    for _ in range(t):
        tree = int(f.readline().strip()) - 1  # 1, 2 => 0, 1
        schedule.append(tree)

    print(max_apple(schedule, t, w))

