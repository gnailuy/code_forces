# http://poj.org/problem?id=1862

from math import sqrt


with open('poj_1862.data', 'r') as f:
    n = int(f.readline())
    stripies = [int(f.readline()) for _ in range(n)]

    while len(stripies) > 1:
        stripies.sort()
        stripies.reverse()
        new_stripy = 2 * sqrt(stripies[0]*stripies[1])
        stripies = stripies[2:] + [new_stripy]

    print('{:.3f}'.format(stripies[0]))

