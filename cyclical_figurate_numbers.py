# -*- coding: utf-8 -*-

def P(i):
    assert i >= 3
    return lambda n: int(n*((i-2)*n+(4-i))/2)


def can_follow(d1, d2):
    assert 1000 <= d1 <= 9999
    assert 1000 <= d2 <= 9999
    return str(d1)[2:] == str(d2)[:2]


def can_append(c, d):  # c: ((...), (...)); d: (n, i)
    seq_d = c[0]
    seq_i = c[1]
    if d[1] in c[1]:
        return False
    else:
        return can_follow(seq_d[-1], d[0])


four_digits = set()
cyclic = set()


for i in range(3, 9):
    n = 1
    while True:
        nxt = P(i)(n)
        if 1000 <= nxt <= 9999:
            four_digits.add((nxt, i))
            cyclic.add(((nxt,), (i,)))
        elif 9999 <= nxt:
            break
        n += 1

for l in range(2, 7):
    new_cyclic = set()

    for c in cyclic:
        # Extend c
        for d in four_digits:
            if can_append(c, d):
                new_cyclic.add((c[0] + (d[0],), c[1] + (d[1],)))

    cyclic = new_cyclic

for c in cyclic:
    if can_follow(c[0][-1], c[0][0]):
        print(c)
        print(sum(c[0]))
        break

