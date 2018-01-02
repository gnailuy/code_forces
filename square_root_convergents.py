#!/usr/local/bin/python3
# -*- coding: utf-8 -*-

def main():
    max_it = 1000
    cnt = 0
    v = (3, 2)

    for _ in range(2, max_it + 1):
        v = (v[1] * 2 + v[0], v[0] + v[1])
        if len(str(v[0])) > len(str(v[1])):
            cnt += 1

    print(cnt)


if __name__ == '__main__':
    main()

