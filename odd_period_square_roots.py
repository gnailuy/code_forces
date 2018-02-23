from math import sqrt


def find_cf(N, a0):
    # Ref: http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/cfINTRO.html#section6.2.2
    # sqrt(N) = a0 + 1 / 1 / (sqrt(N) - a0)
    # x = 1 / (sqrt(N) - a0)
    tail = []

    pam = 1
    div = 1
    a = a0
    # x = div / (pam * (sqrt(N) - a))
    while True:
        new_div = N - a * a
        assert((new_div*pam) % div == 0)
        new_div = int(new_div*pam/div)

        ai = int((sqrt(N)+a)/new_div)
        tail.append(ai)

        a = ai * new_div - a
        div = new_div
    
        if a == a0 and div == 1:
            break

    return a0, tail


upper = 10000
odd_count = 0

for N in range(2, upper+1):
    a0 = int(sqrt(N))
    if a0 ** 2 == N:
        # Ignore perfect square root
        continue

    cf = find_cf(N, a0)
    if len(cf[1]) % 2 != 0:
        odd_count += 1
    print(N, cf)

print(odd_count)

