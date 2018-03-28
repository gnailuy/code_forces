# http://poj.org/problem?id=1017

three_for_two = [0, 5, 3, 1]


with open('poj_1017.data', 'r') as f:
    while True:
        order = [int(x) for x in f.readline().strip().split(' ')]
        if sum(order) == 0:
            break

        parcel = 0
        parcel += order[5] + order[4] + order[3]
        parcel += (order[2] + 3) // 4
        space_for_two = order[3]*5 + three_for_two[order[2]%4]
        if order[1] > space_for_two:
            parcel += (order[1] - space_for_two + 8) // 9
        space_for_one = parcel*36 - order[5]*36 - order[4]*25 - order[3]*16 - order[2]*9 - order[1]*4
        if order[0] > space_for_one:
            parcel += (order[0] - space_for_one + 35) // 36

        print(parcel)

