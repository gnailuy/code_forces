#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m, dsps = 0;
    scanf ("%d %d", &n, &m);

    int windows[2];
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            scanf ("%d %d", &windows[0], &windows[1]);
            if (windows[0] + windows[1] > 0) {
                dsps++;
            }
        }
    }

    printf ("%d\n", dsps);

    return 0;
}

