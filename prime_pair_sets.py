# -*- coding: utf-8 -*-

from miler_rabin_primality import is_prime

prime_cache = {2, 3, 5, 7, 11}
pair_cache = {'3,7'}

four_list = []
three_list = []
two_list = []


def _is_prime(n):
    if n in prime_cache:
        return True
    elif is_prime(n):
        prime_cache.add(n)
        return True
    return False


def _can_pair(n1, n2):  # n1, n2 are both primes
    if n1 >= n2:
        n1, n2 = n2, n1
    str_rep = ','.join((str(n1), str(n2)))
    if str_rep in pair_cache:
        return True

    k1 = int(str(n1) + str(n2))
    k2 = int(str(n2) + str(n1))
    if _is_prime(k1) and _is_prime(k2):
        pair_cache.add(str_rep)
        return True
    else:
        return False


p = 2
while True:
    # Try find `five` set
    found = False
    for s in four_list:
        ok = True
        for si in s:
            if not _can_pair(si, p):
                ok = False
                break
        if ok:
            found = True
            print(s, p, sum(s) + p)
    if found:
        break

    # Try build `four` set
    for s in three_list:
        ok = True
        for si in s:
            if not _can_pair(si, p):
                ok = False
                break
        if ok:
            ns = s.copy()
            ns.add(p)
            four_list.append(ns)

    # Try build `three` set
    for s in two_list:
        ok = True
        for si in s:
            if not _can_pair(si, p):
                ok = False
                break
        if ok:
            ns = s.copy()
            ns.add(p)
            three_list.append(ns)

    # Try build `two` set
    for n in range(2, p):
        if _is_prime(n) and _can_pair(n, p):
            two_list.append({n, p})

    np = p + 1
    while not _is_prime(np):
        np += 1
    p = np  # It's a joke

    print(len(two_list), len(three_list), len(four_list))

