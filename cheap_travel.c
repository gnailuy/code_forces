#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m, a, b;
    scanf ("%d %d %d %d", &n, &m, &a, &b);

    int result = 0, r;
    if (b / (m + 0.0) < a) {
        result += b * (n / m);
        r = a * (n % m);
        result += r < b ? r : b;
    } else {
        result = a * n;
    }

    printf ("%d\n", result);

    return 0;
}

