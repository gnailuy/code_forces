# http://poj.org/problem?id=3040

with open('poj_3040.data', 'r') as f:
    n, c = [int(x) for x in f.readline().strip().split()]
    v_list = []
    v_map = {}
    for _ in range(n):
        v, b = [int(x) for x in f.readline().strip().split()]
        v_list.append(v)
        v_map[v] = b

    v_list.sort()
    v_list.reverse()

    count = 0
    for v in v_list:
        b = v_map[v]
        if v >= c:
            count += b
            del v_map[v]
        else:
            while v in v_map and v_map[v] > 0:
                v_map[v] -= 1
                r = c - v
                for i in range(len(v_list), 0, -1):
                    nv = v_list[i-1]
                    nb = r // nv
                    nb += 1 if r%nv!=0 else 0
                    if v_map[nv] >= nb:
                        v_map[nv] -= nb
                    else:
                        nb = v_map[nv]
                        v_map[nv] = 0
                    r -= nb*nv
                    if r <= 0:
                        break
                if r <= 0:
                    count += 1

    print(count)

