# http://poj.org/problem?id=2376

def read_cows(f, N):
    for _ in range(N):
        s, e = [int(x) for x in f.readline().strip().split(' ')]
        yield (s, e)


def find_max_cow(cows, t, last_cow):
    max_time = -1  # Max time we can go
    cow_index = None  # With this cow

    for i in range(last_cow, len(cows)):
        if cows[i][0] <= t:
            if cows[i][1] > max_time:
                max_time = cows[i][1]
                cow_index = i

    return cow_index


with open('poj_2376.data', 'r') as f:
    while True:
        N, T = [int(x) for x in f.readline().strip().split(' ')]
        if N == 0 or T == 0:
            break

        cows = sorted([x for x in read_cows(f, N)],
                      key=lambda x: x[0])
        t = 1  # Time range start
        last_cow = 0
        count = 0
        while True:
            c = find_max_cow(cows, t, last_cow)
            if c is not None:
                last_cow = c + 1
                count += 1
                t = cows[c][1]
            else:
                break

        if t == T:
            print(count)
        else:
            print(-1)

