# http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0033

def judge(balls):
    tops = [-1, -1]
    for b in balls:
        ok = False
        for i in range(len(tops)):
            if tops[i] < b:
                tops[i] = b
                tops.sort(reverse=True)
                ok = True
                break
        if not ok:
            return 'NO'

    return 'YES'


with open('aoj_0033.data', 'r') as f:
    n = int(f.readline())
    for _ in range(n):
        balls = [int(x) for x in f.readline().split(' ')]
        print(judge(balls))

