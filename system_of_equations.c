#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    int count = 0;
    for (int a = 0; a * a <= n; a++) {
        int b = n - a * a;
        if (a + b * b == m) {
            count++;
        }
    }
    printf ("%d\n", count);

    return 0;
}
