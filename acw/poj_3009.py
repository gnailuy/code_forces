# http://poj.org/problem?id=3009

import copy

def read_board(f, w, h):
    board = []
    for i in range(h):
        line = f.readline().strip().split(' ')
        b_line = [int(x) for x in line]
        if 2 in b_line:
            j = b_line.index(2)
            b_line[j] = 0
            start = (j, i)
        board.append(b_line)
    return board, start


def throw(direction, board, start, w, h):
    x, y = start
    if 'up' == direction:
        if y < 1 or board[y-1][x] == 1:
            return None, None, False
        while y >= 0:
            y -= 1
            if y < 0:
                return None, None, False
            elif board[y][x] == 3:
                return x, y, True
            elif board[y][x] == 1:
                board[y][x] = 0
                return x, y+1, True
    elif 'left' == direction:
        if x < 1 or board[y][x-1] == 1:
            return None, None, False
        while x >= 0:
            x -= 1
            if x < 0:
                return None, None, False
            elif board[y][x] == 3:
                return x, y, True
            elif board[y][x] == 1:
                board[y][x] = 0
                return x+1, y, True
    elif 'down' == direction:
        if y + 1 >= h or board[y+1][x] == 1:
            return None, None, False
        while y < h:
            y += 1
            if y >= h:
                return None, None, False
            elif board[y][x] == 3:
                return x, y, True
            elif board[y][x] == 1:
                board[y][x] = 0
                return x, y-1, True
    elif 'right' == direction:
        if x + 1 >= w or board[y][x+1] == 1:
            return None, None, False
        while x < w:
            x += 1
            if x >= w:
                return None, None, False
            elif board[y][x] == 3:
                return x, y, True
            elif board[y][x] == 1:
                board[y][x] = 0
                return x-1, y, True


def find_solution(board, start, w, h, step=1):
    if step > 10:
        return None, False

    solutions = []
    for direction in ['up', 'down', 'left', 'right']:
        b = copy.deepcopy(board)
        x, y, ok = throw(direction, b, start, w, h)
        if ok:
            if b[y][x] == 3:
                solutions.append(step)
            else:
                result, found = find_solution(b, (x, y), w, h, step=step+1)
                if found:
                    solutions.append(result)

    if len(solutions) > 0:
        return min(solutions), True
    else:
        return None, False


with open('poj_3009.data', 'r') as f:
    while True:
        l = f.readline().strip().split(' ')
        wh = [int(x) for x in l]
        if wh == [0, 0]:
            break

        board, start = read_board(f, wh[0], wh[1])
        solution, ok = find_solution(board, start, wh[0], wh[1])
        if ok:
            print(solution)
        else:
            print(-1)

