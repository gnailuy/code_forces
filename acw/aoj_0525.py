# http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0525

import copy

def read_cracker(f, h, w):
    cracker = []
    for j in range(h):
        cracker.append([int(x) for x in f.readline().strip().split(' ')])
    return cracker


def lead_zero(s, l):
    return '0' * (l - len(s)) + s


def binary_seq(n):
    upper = (2**n)
    l = len(bin(upper-1)) - 2
    for i in range(upper):
        yield lead_zero(bin(i)[2:], l)


def flip_one(i):
    return (i+1)%2


def flip(cracker, seq, h, w):
    for i in range(h):
        if seq[i] == '1':
            for j in range(w):
                cracker[i][j] = flip_one(cracker[i][j])

    for j in range(w):
        total = sum([cracker[i][j] for i in range(h)])
        if total > h//2:
            for i in range(h):
                cracker[i][j] = flip_one(cracker[i][j])

    return cracker


def find_max_shipped(cracker, h, w):
    total = h * w
    burned = None

    for seq in binary_seq(h):
        c = copy.deepcopy(cracker)
        c = flip(c, seq, h, w)
        b = sum([sum(x) for x in c])
        if burned is None or b < burned:
            burned = b

    return total - burned


with open('aoj_0525.data', 'r') as f:
    while True:
        l = f.readline().strip().split(' ')
        h, w = [int(x) for x in l]
        if 0 == h or 0 == w:
            break

        cracker = read_cracker(f, h, w)
        print(find_max_shipped(cracker, h, w))

