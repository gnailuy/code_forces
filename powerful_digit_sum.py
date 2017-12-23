#!/usr/local/bin/python3
# -*- coding: utf-8 -*-

def digit_sum(num):
    ns = str(num)
    return sum([int(x) for x in ns])


def main():
    max_ab = 100
    max_digit_sum = 0
    max_digit_sum_ab = None

    for a in range(1, max_ab):
        for b in range(1, max_ab):
            a_pow_b = a ** b
            ds = digit_sum(a_pow_b)
            if ds > max_digit_sum:
                max_digit_sum = ds
                max_digit_sum_ab = (a, b, a_pow_b)

    print(max_digit_sum)
    print(max_digit_sum_ab)

if __name__ == '__main__':
    main()

