# -*- coding: utf-8 -*-

value_map = {
    'T': 10,
    'J': 11,
    'Q': 12,
    'K': 13,
    'A': 14,
}


def royal_flush(v, s):
    royal = [10, 11, 12, 13, 14]
    if not flush(v, s)[0]:
        return False, None
    else:
        if sorted(v) == royal:
            return True, royal
        else:
            return False, None


def straight_flush(v, s):
    if not flush(v, s)[0] or not straight(v, s)[0]:
        return False, None
    else:
        return True, sorted(v)


def four_of_a_kind(v, s):
    if len(set(v)) != 2:
        return False, None
    i1, i2 = list(set(v))
    if v.count(i1) == 4:
        return True, [i2, i1]
    elif v.count(i2) == 4:
        return True, [i1, i2]
    else:
        return False, None


def full_house(v, s):
    if len(set(v)) != 2:
        return False, None
    i1, i2 = list(set(v))
    if v.count(i1) == 3 and v.count(i2) == 2:
        return True, [i2, i1]
    elif v.count(i2) == 3 and v.count(i1) == 2:
        return True, [i1, i2]
    else:
        return False, None


def flush(v, s):
    return len(set(s)) == 1, sorted(v)


def straight(v, s):
    sv = sorted(v)
    first = sv[0]
    std = [0] + [x - first for x in sv[1:]]
    if std == [0, 1, 2, 3, 4]:
        return True, sorted(v)
    else:
        return False, None


def three_of_a_kind(v, s):
    if len(set(v)) != 3:  # Assume not full house
        return False, None
    i3 = None
    for i in v:
        if v.count(i) == 3:
            i3 = i
    i12 = [x for x in v if x != i3]
    if i3 is not None:
        return True, sorted(i12) + [i3]
    return False, None


def two_pairs(v, s):
    if len(set(v)) != 3:
        return False, None
    res = []
    for i in v:
        if v.count(i) == 2 and i not in res:
            res.append(i)
    if len(res) == 2:
        i1 = [x for x in v if x not in res]
        return True, i1 + sorted(res)
    else:
        return False, None


def one_pair(v, s):
    if len(set(v)) != 4:
        return False, None
    i2 = None
    for i in v:
        if v.count(i) == 2:
            i2 = i
    if i2 is not None:
        i1 = [x for x in v if x != i2]
        return True, sorted(i1) + [i2]
    else:
        return False, None


rule_list = [
    royal_flush,
    straight_flush,
    four_of_a_kind,
    full_house,
    flush,
    straight,
    three_of_a_kind,
    two_pairs,
    one_pair,
]


def rank_hand(hand):
    values = [value_map[y] if y in value_map else int(y)
                 for y in [x[0] for x in hand]]
    suits = [x[1] for x in hand]
    for i, rule in enumerate(rule_list):
        ok, v = rule(values, suits)
        if ok:
            return i, v
    return len(rule_list), sorted(values)

p1_win_count = 0

with open('p054_poker.txt') as f:
    for l in f:
        l = l.strip()
        p = l.split(' ')
        p1 = p[0:5]
        p2 = p[5:]

        print 'Working on', l

        r1, v1 = rank_hand(p1)
        r2, v2 = rank_hand(p2)
        if r1 < r2:
            p1_win_count += 1
        elif r1 == r2:
            assert len(v1) == len(v2)
            vv1 = [x for x in reversed(v1)]
            vv2 = [x for x in reversed(v2)]
            for i, v in enumerate(vv1):
                if v > vv2[i]:
                    print 'P1 Win'
                    p1_win_count += 1
                    break
                elif v < vv2[i]:
                    break

print p1_win_count

