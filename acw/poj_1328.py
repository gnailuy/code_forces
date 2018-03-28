# http://poj.org/problem?id=1328

import copy
import math

def calculate_bound(x, y, d):
    if y > d:
        return None
    else:
        delta = math.sqrt(d*d - y*y)
        return (x - delta, x + delta)


def read_island(f, n, d):
    for _ in range(n):
        x, y = [int(x) for x in f.readline().strip().split(' ')]
        yield calculate_bound(x, y, d)
    f.readline()


def minimal_radar(islands):
    radar = 0
    il = copy.deepcopy(islands)
    while True:
        if len(il) == 0:
            break

        station = il[0][1]  # 尽左安装一个雷达，覆盖右脚最左的岛屿
        radar += 1

        # 删掉此雷达覆盖的所有岛屿
        to_delete = [x for x in il if x[0] <= station <= x[1]]
        il = [x for x in il if x not in to_delete]

    return radar


with open('poj_1328.data', 'r') as f:
    i = 1
    while True:
        n, d = [int(x) for x in f.readline().strip().split(' ')]
        if n == 0 and d == 0:
            break

        islands = [x for x in read_island(f, n, d)]
        if None in islands:
            print('Case {}: -1'.format(i))
        else:
            islands.sort(key=lambda x: x[1])
            m = minimal_radar(islands)
            print('Case {}: {}'.format(i, m))
        i += 1

