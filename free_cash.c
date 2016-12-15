#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, min_cashes = 0;
    scanf ("%d", &n);

    int cashes = 1, h, m, in_time, last_in_time = -1;
    for (int i = 0; i < n; i++) {
        scanf ("%d %d", &h, &m);
        in_time = h * 60 + m;
        if (in_time == last_in_time) {
            cashes += 1;
        } else {
            if (cashes > min_cashes) {
                min_cashes = cashes;
            }
            cashes = 1;
        }
        last_in_time = in_time;
    }
    if (cashes > min_cashes) {
        min_cashes = cashes;
    }

    printf ("%d\n", min_cashes);

    return 0;
}

