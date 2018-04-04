# http://poj.org/problem?id=3616

def solve(n, m, r, schedule):
    max_array = [0 for _ in range(m)]
    max_array[0] = schedule[0][2]

    for i in range(1, m):
        local_max = []
        si = schedule[i]
        for j in range(i):
            sj = schedule[j]
            if si[0] - sj[1] >= r:
                local_max.append(max_array[j] + si[2])
        max_array[i] = max(local_max) if len(local_max) > 0 else si[2]

    return max(max_array)


with open('poj_3616.data', 'r') as f:
    n, m, r = [int(x) for x in f.readline().strip().split(' ')]

    schedule = []
    for _ in range(m):
        start, end, efficiency = [int(x) for x in f.readline().strip().split(' ')]
        schedule.append((start, end, efficiency))

    schedule.sort(key=lambda x: x[1])  # Sort by end time

    print(solve(n, m, r, schedule))

