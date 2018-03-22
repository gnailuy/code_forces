# http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0558

def read_town(f, h, w, n):
    board = []
    factory_list = [0] * (n+1)

    for i in range(h):
        l = [x for x in f.readline().strip()]
        if 'S' in l:
            j = l.index('S')
            factory_list[0] = (i, j)
        for k in range(w):
            c = l[k]
            if c not in ['S', 'X', '.']:
                factory_list[int(c)] = (i, k)
        board.append(l)

    return board, factory_list


def find_neighbors(board, h, w, start, visited):
    neighbors = []

    i, j = start
    if j > 0 and board[i][j-1] != 'X':
        neighbors.append((i, j-1))
    if j + 1 < w and board[i][j+1] != 'X':
        neighbors.append((i, j+1))
    if i > 0 and board[i-1][j] != 'X':
        neighbors.append((i-1, j))
    if i + 1 < h and board[i+1][j] != 'X':
        neighbors.append((i+1, j))

    return [x for x in neighbors if x not in visited]


def path_length(board, h, w, factory_list, i):
    step = 0
    end = factory_list[i+1]
    to_visit = [factory_list[i]]
    visited = []
    while True:
        if end in to_visit:
            break
        else:
            new_to_visit = []
            visited += to_visit
            for s in to_visit:
                new_to_visit += find_neighbors(board, h, w, s, visited)
            to_visit = new_to_visit
            step += 1

    return step


with open('aoj_0558.data', 'r') as f:
    while True:
        l = f.readline().strip().split(' ')
        hwn = [int(x) for x in l]
        if hwn == [0, 0, 0]:
            break
        board, factory_list = read_town(f, hwn[0], hwn[1], hwn[2])
        total = 0
        for i in range(hwn[2]):
            total += path_length(board, hwn[0], hwn[1], factory_list, i)
        print(total)

