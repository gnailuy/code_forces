#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <math.h>

int main (int argc, char * argv[]) {
    int a, b, n;
    scanf ("%d %d", &a, &b);
    scanf ("%d", &n);

    int x, y, v;
    double distance, time, min_time = -1;
    for (int i = 0; i < n; i++) {
        scanf ("%d %d %d", &x, &y, &v);
        distance = sqrt ((((double) x) - a) * (x - a) + (y - b) * (y - b));
        time = distance / v;
        if (time < min_time || -1 == min_time) {
            min_time = time;
        }
    }
    printf ("%0.20f\n", min_time);

    return 0;
}

