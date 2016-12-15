#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, d, song_duration, sum_duration = 0;
    scanf ("%d %d", &n, &d);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &song_duration);
        sum_duration += song_duration;
    }

    if (sum_duration + (n - 1) * 10 > d) {
        printf ("-1\n");
    } else {
        printf ("%d\n",
                2 * (n - 1) + (d - sum_duration - (n - 1) * 10) / 5);
    }

    return 0;
}

