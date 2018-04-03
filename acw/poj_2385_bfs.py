# http://poj.org/problem?id=2385

def max_apple(schedule, w):
    point = [(1, 0, 0)]  # (position, moves, apples)
    for s in schedule:
        new_point = []
        for p in point:
            new_point.append((p[0], p[1], p[2] if p[0] != s else p[2]+1))  # Don't move
            if p[1] < w:
                new_point.append((1 if p[0] == 2 else 2, p[1]+1, p[2] if p[0] == s else p[2]+1))  # Move
        point = new_point

    max_apple = 0
    for p in point:
        if max_apple < p[2]:
            max_apple = p[2]
    return max_apple


with open('poj_2385.data', 'r') as f:
    t, w = [int(x) for x in f.readline().strip().split(' ')]
    schedule = []
    for _ in range(t):
        tree = int(f.readline().strip())
        schedule.append(tree)

    print(max_apple(schedule, w))

