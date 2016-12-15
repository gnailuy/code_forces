#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

long long nC2 (int n) {
    return ((long long) n) * (n - 1) / 2;
}

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    long long min = ((long long) m) * nC2 (n / m) + (n % m) * (n / m);
    long long max = nC2 (n - m + 1);

    printf ("%I64d %I64d\n", min, max);

    return 0;
}

