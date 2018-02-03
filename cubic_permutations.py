# -*- coding: utf-8 -*-

def cubes_with_x_digits(x):
    from math import ceil, floor
    lower = floor((10 ** (x - 1)) ** (1. / 3))
    upper = ceil((10 ** x - 1) ** (1. / 3))
    for i in range(lower, upper):
        yield i ** 3


def sort_digits(n):
    return ''.join([str(x) for x in sorted([int(c) for c in str(n)])])


digits = 3
while True:
    cubes = [x for x in cubes_with_x_digits(digits)]
    cube_repr = [sort_digits(x) for x in cubes]
    repr_cnt = [cube_repr.count(x) >= 5 for x in cube_repr]
    if any(repr_cnt):
        print([x[0] for x in zip(cubes, repr_cnt) if x[1]])
        break
    digits += 1

