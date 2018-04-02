# http://poj.org/problem?id=2229

def sumsets(n):
    store = [0 for _ in range(n+1)]
    store[1] = 1
    store[2] = 2

    for i in range(3, n+1):
        if i % 2 != 0:
            store[i] = store[i-1]
        else:
            store[i] = store[i-2] + store[int(i/2)]

    return store[n]


with open('poj_2229.data', 'r') as f:
    n = int(f.readline().strip())
    print(sumsets(n))

