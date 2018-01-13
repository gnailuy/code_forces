# -*- coding: utf-8 -*-

from operator import xor

textable_chars = [' ', '!', '"', '\'', '(', ')', ',', '.', ':', ';', '?']

with open('p059_cipher.txt', 'r') as f:
    l = f.readline().strip()
    chars = [int(x) for x in l.split(',')]

    c_three = [
        [x[1] for x in enumerate(chars) if x[0] % 3 == 0],
        [x[1] for x in enumerate(chars) if x[0] % 3 == 1],
        [x[1] for x in enumerate(chars) if x[0] % 3 == 2],
    ]

    p_three = [None, None, None]

    for i in range(3):
        c = c_three[i]
        for k in range(ord('a'), ord('z') + 1):
            c_descrypted = [xor(ch, k) for ch in c]
            textable = [x in [ord(t) for t in textable_chars] \
                            or x in [t for t in range(ord('0'), ord('9') + 1)] \
                            or x in [t for t in range(ord('a'), ord('z') + 1)] \
                            or x in [t for t in range(ord('A'), ord('Z') + 1)]
                        for x in c_descrypted]
            if all(textable):
                p_three[i] = k

    text_code = []
    for i, c in enumerate(chars):
        text_code.append(xor(c, p_three[i%3]))
    text = [chr(x) for x in text_code]
    password = [chr(x) for x in p_three]
    print(''.join(password))
    print(''.join(text))
    print(sum(text_code))

