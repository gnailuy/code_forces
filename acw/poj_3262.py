# http://poj.org/problem?id=3262

with open('poj_3262.data', 'r') as f:
    n = int(f.readline())
    t_list = []
    d_list = []
    for _ in range(n):
        t, d = [int(x) for x in f.readline().strip().split(' ')]
        t_list.append(t)
        d_list.append(d)

    total_flower = 0
    while len(t_list) > 0:
        min_flower = None
        choice = None
        for i in range(len(t_list)):
            ti = t_list[i] * 2
            d_sum = sum(d_list) - d_list[i]
            flower = ti * d_sum
            if min_flower is None or flower < min_flower:
                min_flower = flower
                choice = i

        total_flower += min_flower
        t_list = t_list[0:choice] + t_list[choice+1:]
        d_list = d_list[0:choice] + d_list[choice+1:]

    print(total_flower)

