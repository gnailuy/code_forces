#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, percent;
    double sum_percent = 0.0;
    scanf ("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &percent);
        sum_percent += percent;
    }

    printf ("%0.12f\n", sum_percent / n);

    return 0;
}

