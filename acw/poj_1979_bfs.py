# http://poj.org/problem?id=1979

def read_block(fp, w, h):
    block = []
    ix = -1
    iy = -1

    for i in range(h):
        l = fp.readline()
        bl = []
        for j in range(w):
            if l[j] == '.':
                bl.append(0)  # 0 for black
            elif l[j] == '#':
                bl.append(1)  # 1 for red
            elif l[j] == '@':
                bl.append(0)
                ix, iy = j, i
            else:  # Error
                return None, None
        block.append(bl)

    return block, (ix, iy)


def go(direction, initial, w, h):
    x, y = initial[0], initial[1]
    if 'up' == direction:
        y -= 1
        if y >= 0:
            return x, y
    elif 'down' == direction:
        y += 1
        if y < h:
            return x, y
    elif 'left' == direction:
        x -= 1
        if x >= 0:
            return x, y
    elif 'right' == direction:
        x += 1
        if x < w:
            return x, y

    return None, None


def iterate_path(block, initial, w, h):
    block[initial[1]][initial[0]] = 2  # Mark initial point
    count = 1  # Count for the initial point

    to_visit = [(initial)]
    while len(to_visit) > 0:
        new_to_visit = []
        for p in to_visit:
            for direction in ['up', 'down', 'left', 'right']:  # Go to four directions
                x, y = go(direction, p, w, h)
                if x is not None and block[y][x] == 0:
                    block[y][x] = 2  # Mark the new point
                    count += 1  # Count fot the new point
                    new_to_visit.append((x, y))
        to_visit = new_to_visit

    return count


with open('poj_1979.data', 'r') as f:
    while True:
        l = f.readline()
        wh = [int(x) for x in l.strip().split(' ')]
        if 0 == wh[0] and 0 == wh[1]:
            break
        else:
            block, initial = read_block(f, wh[0], wh[1])
            print(iterate_path(block, initial, wh[0], wh[1]))

