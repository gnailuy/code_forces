# -*- coding: utf-8 -*-

import sys

upper_n = 10000
pn_map = {}
pn_list = []


for n in xrange(1, upper_n):
    pn = n * (3 * n - 1) / 2
    pn_map[pn] = n
    pn_list.append(pn)


for i in xrange(0, upper_n-1):
    for j in xrange(i+1, upper_n-1):
        pn_sum = pn_list[j] + pn_list[i]
        pn_diff = pn_list[j] - pn_list[i]
        if pn_sum in pn_map and pn_diff in pn_map:
            print pn_map[pn_list[i]], pn_list[i]
            print pn_map[pn_list[j]], pn_list[j]
            print pn_map[pn_sum], pn_sum
            print pn_map[pn_diff], pn_diff
            sys.exit(0)

