# -*- coding: utf-8 -*-

prime_list = [2, 3, 5, 7, 11, 13, 17]

def all_perms(elements):
    if len(elements) <=1:
        yield elements
    else:
        for perm in all_perms(elements[1:]):
            for i in range(len(elements)):
                yield perm[:i] + elements[0:1] + perm[i:]

i = 123456789
i_str = '{:010d}'.format(i)

total = 0
for p in all_perms(i_str):
    ok = True
    for i in range(0, len(prime_list)):
        if int(p[i+1:i+4]) % prime_list[i] != 0:
            ok = False
            break
    if ok:
        total += int(p)
        print p

print total

