# http://poj.org/problem?id=3050

def read_scotch(f, h, w):
    scotch = []
    for j in range(h):
        l = f.readline().strip().split(' ')
        scotch.append([int(x) for x in l])
    return scotch


def index_to_point(i, h, w):  # (line, column)
    return (i//h, i%w)


def point_to_index(p, h, w):  # (line, column)
    return p[0]*h + p[1]


def generate_all_step(posibilities, steps):
    for p in range(posibilities):
        if steps == 1:
            yield [p]
        else:
            for s in generate_all_step(posibilities, steps-1):
                yield [p] + s


def get_path(scotch, si, sj, h, w, step):
    i, j = si, sj
    path = [scotch[i][j]]
    for s in step:
        if 0 == s:  # up
            if i > 0:
                i -= 1
                path.append(scotch[i][j])
            else:
                return None, False
        elif 1 == s:  # down
            if i < h-1:
                i += 1
                path.append(scotch[i][j])
            else:
                return None, False
        elif 2 == s:  # left
            if j > 0:
                j -= 1
                path.append(scotch[i][j])
            else:
                return None, False
        elif 3 == s:  # right
            if j < w-1:
                j += 1
                path.append(scotch[i][j])
            else:
                return None, False

    return path, True


def generate_all_path(scotch, start, h, w):
    i, j = start
    steps = generate_all_step(4, 5)
    for step in steps:
        path, ok = get_path(scotch, i, j, h, w, step)
        if ok:
            yield path


with open('poj_3050.data', 'r') as f:
    while True:
        l = f.readline().strip().split(' ')
        h, w = [int(x) for x in l]
        if 0 == h or 0 == w:
            break
        scotch = read_scotch(f, h, w)

        path_set = set()
        for i in range(h*w):
            for p in generate_all_path(scotch, index_to_point(i, h, w), h, w):
                path_set.add(''.join([str(x) for x in p]))
        print(len(path_set))

