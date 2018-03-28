# http://poj.org/problem?id=3190

with open('poj_3190.data', 'r') as f:
    n = int(f.readline())
    cows = []
    for i in range(n):
        s, e = [int(x) for x in f.readline().strip().split(' ')]
        cows.append((s, e, i))

    cows.sort(key=lambda x: x[0])

    stall_index = 0
    spare_stall = set()
    stall_endtime_list = []
    history = []

    for c in cows:
        now = c[0]
        to_free = [x for x in stall_endtime_list if x[1] < now]
        for t in to_free:
            spare_stall.add(t[0])
        stall_endtime_list = [x for x in stall_endtime_list if x not in to_free]

        if len(spare_stall) == 0:
            stall_index += 1  # Add a stall
            stall_endtime_list.append((stall_index, c[1]))
            history.append((c[2], stall_index))
        else:
            stall = spare_stall.pop()
            stall_endtime_list.append((stall, c[1]))
            history.append((c[2], stall))

        stall_endtime_list.sort(key=lambda x: x[1])

    print(stall_index)
    history.sort(key=lambda x: x[0])
    for h in history:
        print(h[1])

