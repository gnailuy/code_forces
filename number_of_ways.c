#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    long long sum = 0, one_third;
    int * a = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &a[i]);
        sum += a[i];
    }
    int result = 0;

    if (sum % 3 == 0) {
        one_third = sum / 3;

        long long first = 0;
        for (int i = 0; i < n - 2; i++) {
            first += a[i];
            if (one_third == first) {
                long long second = 0;
                for (int j = i + 1; j < n - 1; j++) {
                    second += a[j];
                    if (one_third == second) {
                        result++;
                    }
                }
            }
        }
    }

    printf ("%d\n", result);

    free (a);
    return 0;
}

