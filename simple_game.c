#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m, a;
    scanf ("%d %d", &n, &m);

    int middle = n / 2 + n % 2;
    if (m < middle) {
        a = m + 1;
    } else if (m > middle) {
        a = m - 1;
    } else {
        a = m + (n == 1 ? 0 : (n % 2 == 0 ? 1 : -1));
    }

    printf ("%d\n", a);

    return 0;
}

