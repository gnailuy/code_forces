# http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0118

def read_block(fp, h, w):
    block = []

    for i in range(h):
        l = fp.readline().strip()

        bl = []
        for j in range(w):
            if l[j] == '@':
                bl.append('a')  # a for apple
            elif l[j] == '#':
                bl.append('y')  # y for oyster
            elif l[j] == '*':
                bl.append('o')  # o for orange
            else:  # Error
                return None
        block.append(bl)

    return block


def find_start(block, h, w):
    for i in range(h):
        for j in range(w):
            if block[i][j] in ['a', 'y', 'o']:
                return (i, j)


def mark_area(block, h, w, start, i):
    fruit = block[start[0]][start[1]]

    to_visit = [start]
    block[start[0]][start[1]] = i
    while len(to_visit) > 0:
        new_to_visit = []
        for p in to_visit:
            x, y = p

            ux, uy = x, y-1
            if uy >= 0 and fruit == block[ux][uy]:
                block[ux][uy] = i
                new_to_visit.append((ux, uy))

            dx, dy = x, y+1
            if dy < h and fruit == block[dx][dy]:
                block[dx][dy] = i
                new_to_visit.append((dx, dy))

            lx, ly = x-1, y
            if lx >= 0 and fruit == block[lx][ly]:
                block[lx][ly] = i
                new_to_visit.append((lx, ly))

            rx, ry = x+1, y
            if rx < w and fruit == block[rx][ry]:
                block[rx][ry] = i
                new_to_visit.append((rx, ry))

        to_visit = new_to_visit


def devide_orchard(block, h, w):
    i = 0
    while True:
        start = find_start(block, h, w)
        if start is None:
            break
        mark_area(block, h, w, start, str(i))
        i += 1

    return i


with open('aoi_0118.data', 'r') as f:
    while True:
        l = f.readline()
        hw = [int(x) for x in l.strip().split(' ')]
        if 0 == hw[0] and 0 == hw[1]:
            break
        else:
            block = read_block(f, hw[0], hw[1])
            count = devide_orchard(block, hw[0], hw[1])
            print(count)

