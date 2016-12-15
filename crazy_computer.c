#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, c, t;
    scanf ("%d %d", &n, &c);

    int words = 0, last_t = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &t);
        if (t - last_t > c) {
            words = 1;
        } else {
            words++;
        }
        last_t = t;
    }

    printf ("%d\n", words);

    return 0;
}

