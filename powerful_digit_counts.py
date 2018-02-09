max_n = 1000

n = 1
while True:
    i = 1
    met = False
    while True:
        p = (i ** n)
        p_len = len(str(p))
        if len(str(p)) == n:
            met = True
            print(p, i, n)
        if len(str(p)) > n:
            break
        i += 1
    if not met:
        break
    n += 1

