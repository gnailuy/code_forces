#!/usr/local/bin/python3
# -*- coding: utf-8 -*-

from miler_rabin_primality import is_prime


def main():
    total_count = 1
    prime_count = 0

    i = 1
    while True:
        i += 2
        total_count += 4

        bottom_right = i * i
        for j in range(0, 4):
            x = bottom_right - (i-1) * j
            if is_prime(x):
                prime_count += 1

        ratio = prime_count / total_count
        print(i, prime_count, total_count, ratio)
        if ratio <= 0.1:
            break


if __name__ == '__main__':
    main()

