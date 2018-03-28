# http://poj.org/problem?id=2393

with open('poj_2393.data', 'r') as f:
    n, s = [int(x) for x in f.readline().strip().split(' ')]
    schedule = []
    total_cost = 0

    for i in range(n):
        c, y = [int(x) for x in f.readline().strip().split(' ')]
        fee_today = c*y
        fee_before = None
        for j in range(len(schedule)):
            cd = schedule[j][0]
            fee_j = cd*y + (i-j)*y*s
            if fee_before is None or fee_before > fee_j:
                fee_before = fee_j

        if fee_before is None or fee_before > fee_today:
            total_cost += fee_today
        else:
            total_cost += fee_before

        schedule.append((c, y))

    print(total_cost)

