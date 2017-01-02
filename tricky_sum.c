#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

long long tricky_sum (int n) {
    long long sum = ((long long) n) * (n + 1) / 2;
    int upper = 1, upper_power = 0;
    while (upper <= n / 2) {
        upper *= 2;
        upper_power++;
    }
    long long sum_of_twos = (1LL << (upper_power + 1)) - 1;
    return sum - sum_of_twos * 2;
}

int main (int argc, char * argv[]) {
    int t, n;
    scanf ("%d", &t);

    for (int i = 0; i < t; i++) {
        scanf ("%d", &n);
        // printf ("%I64d\n", tricky_sum (n));
        printf ("%lld\n", tricky_sum (n));
    }

    return 0;
}

