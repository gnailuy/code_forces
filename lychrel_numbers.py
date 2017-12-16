# -*- coding: utf-8 -*-

upper = 10000

def reverse_num(num):
    return int(str(num)[::-1])


def is_palindromic(num):
    if num == reverse_num(num):
        return True
    return False


def is_lychrel(num):
    for i in xrange(0, 50):
        num = num + reverse_num(num)
        if is_palindromic(num):
            return False
    return True


cnt = 0

for i in xrange(1, upper):
    if is_lychrel(i):
        cnt += 1

print cnt

