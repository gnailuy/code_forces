# http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0121

def switch(state, i, j):
    sl = list(state)
    sl[i], sl[j] = sl[j], sl[i]
    return tuple(sl)


def next_state(state, taboo):
    next_set = set()
    zero = state.index(0)
    i, j = zero%4, zero//4

    if i > 0:
        neighbor = i - 1 + j*4
        next_set.add(switch(state, zero, neighbor))
    if j > 0:
        neighbor = i + (j-1)*4
        next_set.add(switch(state, zero, neighbor))
    if i < 3:
        neighbor = i + 1 + j*4
        next_set.add(switch(state, zero, neighbor))
    if j < 1:
        neighbor = i + (j+1)*4
        next_set.add(switch(state, zero, neighbor))

    return {x for x in next_set if x not in taboo}


def one_step(css, taboo):
    next_set = set()
    for s in css:
        next_set = next_set.union(next_state(s, taboo))
    return next_set, taboo.union(next_set)


def solve(state):
    final_state = tuple([x for x in range(8)])
    met_state_set = {tuple(state)}
    current_state_set = {tuple(state)}

    step = 0
    while True:
        if final_state in current_state_set:
            return step
        current_state_set, met_state_set = one_step(current_state_set, met_state_set)
        step += 1


with open('aoj_0121.data', 'r') as f:
    while True:
        l = f.readline().strip()
        if not l:
            break

        state = [int(x) for x in l.split(' ')]
        print(solve(state))

