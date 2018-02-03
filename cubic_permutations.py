# -*- coding: utf-8 -*-

from math import ceil, floor


def cubes_of_power(x):
    lower = floor((10 ** (x - 1)) ** (1. / 3))
    upper = ceil((10 ** x - 1) ** (1. / 3))
    for i in range(lower, upper):
        yield i ** 3


def sort_digits(n):
    return ''.join([str(x) for x in sorted([int(c) for c in str(n)])])


digits = 2
while True:
    cubes = [x for x in cubes_of_power(digits)]
    cube_reps = [sort_digits(x) for x in cubes]
    counts = dict()
    for i in cube_reps:
        counts[i] = counts.get(i, 0) + 1

    found = False
    for i in counts:
        if counts[i] == 5:
            print([x for x in cubes if sort_digits(x) == i])
            found = True
    if found:
        break

    digits += 1

