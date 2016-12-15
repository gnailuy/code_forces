#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MASK 0x00000001

int ones_cnt (int bits) {
    int cnt = 0;
    while (bits > 0) {
        if ((bits & MASK) == 1) cnt++;
        bits = bits >> 1;
    }
    return cnt;
}

int main (int argc, char * argv[]) {
    int n, m, k, friends = 0;
    scanf ("%d %d %d", &n, &m, &k);

    int * armies = (int *) malloc (sizeof(int) * (m + 1));
    for (int i = 0; i <= m; i++) {
        scanf ("%d", &armies[i]);
    }

    for (int i = 0; i < m; i++) {
        int diffs = armies[m] ^ armies[i];
        if (ones_cnt (diffs) <= k) {
            friends++;
        }
    }

    printf ("%d\n", friends);

    free (armies);
    return 0;
}

