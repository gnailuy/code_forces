# http://poj.org/problem?id=3187

def all_sequences(seq):
    if len(seq) == 1:
        yield seq

    for i in range(len(seq)):
        for s in all_sequences(seq[:i] + seq[i+1:]):
            yield [seq[i]] + s


def add_down(seq):
    new_seq = seq[:]
    while True:
        n = len(new_seq)
        if 1 == n:
            return new_seq[0]

        next_seq = []
        for i in range(n-1):
            next_seq.append(new_seq[i] + new_seq[i+1])

        new_seq = next_seq


def find_sequence(n, s):
    for seq in all_sequences(list(range(1, n+1))):
        if add_down(seq) == s:
            return ' '.join([str(x) for x in seq])


with open('poj_3187.data', 'r') as f:
    while True:
        l = f.readline().strip()
        if not l:
            break
        n, the_sum = [int(x) for x in l.split(' ')]
        print(find_sequence(n, the_sum))

