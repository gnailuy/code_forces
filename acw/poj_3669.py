# http://poj.org/problem?id=3669

def read_schedule(f, m):
    schedule = []
    for i in range(m):
        l = f.readline().strip().split(' ')
        xyt = tuple([int(x) for x in l])
        schedule.append(xyt)
    schedule.sort(key=lambda x: x[2])
    return schedule


def unsafe_place(schedule):
    unsafe_set = set()
    ticker = 0
    while schedule[-1][2] >= ticker:
        boom(schedule, unsafe_set, ticker)
        ticker += 1
    return unsafe_set


def boom(schedule, taboo_set, ticker):
    for s in schedule:
        if s[2] == ticker:
            x, y = s[0], s[1]
            taboo_set.add((x, y))
            if x > 0:
                taboo_set.add((x-1, y))
            if y > 0:
                taboo_set.add((x, y-1))
            taboo_set.add((x+1, y))
            taboo_set.add((x, y+1))
        elif s[2] > ticker:
            break


def go(p, taboo_set):
    x, y = p
    to_go = []
    if x-1>0:
        to_go.append((x-1, y))
    if y-1>0:
        to_go.append((x, y-1))
    to_go.append((x+1, y))
    to_go.append((x, y+1))
    return [x for x in to_go if x not in taboo_set]


def find_shortest_path(schedule):
    visited = [(0, 0)]
    all_unsafe = unsafe_place(schedule)

    ticker = 0
    taboo_set = set()
    while True:
        # Drop the boom
        boom(schedule, taboo_set, ticker)

        # Check if we were boomed
        visited = [p for p in visited if p not in taboo_set]
        if len(visited) == 0:
            return -1

        # Check if we are in a safe place
        safe_list = [(p not in all_unsafe) for p in visited]
        if any(safe_list):
            return ticker

        # Run
        to_visit = []
        for p in visited:
            to_visit += go(p, taboo_set)
        visited = to_visit

        # Time flies
        ticker += 1


with open('poj_3669.data', 'r') as f:
    while True:
        m = int(f.readline())
        if m == 0:
            break
        schedule = read_schedule(f, m)
        time = find_shortest_path(schedule)
        print(time)

