#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX 1000000000

int main (int argc, char * argv[]) {
    int n, number, odd_count = 0, min_odd = MAX - 1;
    scanf ("%d", &n);

    long long sum = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &number);
        sum += number;
        if (number % 2 != 0) {
            odd_count++;
            if (min_odd > number) {
                min_odd = number;
            }
        }
    }

    if (odd_count % 2 != 0) {
        sum -= min_odd;
    }
    printf ("%lld\n", sum);
    // printf ("%I64d\n", sum);

    return 0;
}

