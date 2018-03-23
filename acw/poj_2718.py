# http://poj.org/problem?id=2718

def all_permutations(l):
    n = len(l)
    if n == 1:
        yield l
    else:
        for i in range(n):
            for sub in all_permutations(l[:i] + l[i+1:]):
                yield [l[i]] + sub


def zero_lead(s):
    return len(s) > 1 and s[0] == '0'


def to_string(l):
    return ''.join([str(x) for x in l])


def abs(n):
    return n if n >= 0 else -n


with open('poj_2718.data', 'r') as f:
    n = int(f.readline())

    for _ in range(n):
        min_diff = None
        l = [int(x) for x in f.readline().strip().split(' ')]

        d2 = len(l)//2
        d1 = len(l) - d2

        for p in all_permutations(l):
            first = to_string(p[:d1])
            second = to_string(p[d1:])
            if zero_lead(first) or zero_lead(second):
                continue

            diff = abs(int(first) - int(second))
            if min_diff is None or diff < min_diff:
                min_diff = diff

        print(min_diff)

