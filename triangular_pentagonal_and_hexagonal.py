# -*- coding: utf-8 -*-

import sys


def triangle(n):
    return n * (n + 1) / 2


def pentagonal(n):
    return n * (3*n - 1) / 2


def hexagonal(n):
    return n * (2*n - 1)


i = 285
pn = 165
hn = 143
pu = pentagonal(pn)
hu = hexagonal(hn)

while True:
    i += 1

    t = triangle(i)
    while t > pu:
        pn += 1
        pu = pentagonal(pn)
    while t > hu:
        hn += 1
        hu = hexagonal(hn)
    if t == pu and t == hu:
        print i, t, pu, hu
        break

