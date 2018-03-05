# Copied from: https://www.mathblog.dk/project-euler-66-diophantine-equation/
# #NEEDFURTHERSTUDY

from math import sqrt

result = 0
xmax = 0

for D in range(2, 1000+1):
    limit = int(sqrt(D))
    if limit * limit == D:
        continue

    m = 0
    d = 1
    a = limit

    numm1 = 1
    num = a

    denm1 = 0
    den = 1

    while num*num - D*den*den != 1:
        m = d * a - m
        d = int((D - m * m) / d)
        a = int((limit + m) / d)

        numm2 = numm1
        numm1 = num
        denm2 = denm1
        denm1 = den

        num = a * numm1 + numm2
        den = a * denm1 + denm2

    if num > xmax:
        xmax = num
        result = D


print(result, xmax)

