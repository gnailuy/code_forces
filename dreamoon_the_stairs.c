#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    if (n < m) {
        printf ("-1\n");
    } else {
        int min_steps = n / 2 + (n % 2 == 0 ? 0 : 1);
        for (int i = min_steps; i <= n; i++) {
            if (i % m == 0) {
                printf ("%d\n", i);
                break;
            }
        }
    }

    return 0;
}

