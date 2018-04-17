# http://poj.org/problem?id=1742

def solve(n, m, a_list, c_map):
    # TODO
    pass


with open('poj_1742.data', 'r') as f:
    while True:
        n, m = [int(x) for x in f.readline().strip().split(' ')]
        if n == 0 and m == 0:
            break

        ac_list = [int(x) for x in f.readline().strip().split(' ')]
        a_list = ac_list[:n]
        c_map = dict([(a_list[i], ac_list[n+i]) for i in range(n)])

        a_list.sort()
        a_list.reverse()

        print(solve(n, m, a_list, c_map))

