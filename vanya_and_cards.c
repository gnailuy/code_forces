#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

inline int abs (int num) {
    return num >= 0 ? num : -num;
}

int main (int argc, char * argv[]) {
    int n, x, a;
    scanf ("%d %d", &n, &x);

    int sum = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &a);
        sum += a;
    }

    sum = abs (sum);
    int cards = sum / x;
    if (sum % x != 0) {
        cards++;
    }
    printf ("%d\n", cards);

    return 0;
}

