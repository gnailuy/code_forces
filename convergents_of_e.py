import sys

def ith_constant(i):
    return int((i-2)/3*2+2) if (i-2)%3 == 0 else 1


def seq_n(n):
    return [ith_constant(x) for x in range(1, n+1)]


def convergents(e, i):
    if len(e) == i + 1:
        return 1, e[i]
    else:
        n, d = convergents(e, i+1)
        return d, e[i]*d+n


def main(argv):
    if len(argv) < 2:
        upper = 99
    else:
        upper = int(argv[1]) - 1
    e = [2] + seq_n(upper)
    d, n = convergents(e, 0)
    print(sum(int(i) for i in str(n)))


if __name__ == '__main__':
    main(sys.argv)

