# http://poj.org/problem?id=3616

def solve(n, m, r, schedule):
    status = [(0, 0)]  # earliest start time, total milk
    for s in schedule:
        new_status = []
        for u in status:
            start = s[0]
            end = s[1]
            efficiency = s[2]
            if u[0] <= start:
                new_status.append((end + r, u[1] + efficiency))  # milk
                new_status.append(u)  # don't milk
            else:
                new_status.append(u)  # can't milk

        status = new_status

    milk = [x[1] for x in status]
    return max(milk)


with open('poj_3616.data', 'r') as f:
    n, m, r = [int(x) for x in f.readline().strip().split(' ')]

    schedule = []
    for _ in range(m):
        start, end, efficiency = [int(x) for x in f.readline().strip().split(' ')]
        schedule.append((start, end, efficiency))

    schedule.sort(key=lambda x: x[0])  # Sort by start time

    print(solve(n, m, r, schedule))

