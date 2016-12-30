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
    long long result = 0;

    if (sum % 3 == 0) {
        one_third = sum / 3;

        int * tail = (int *) malloc (sizeof(int) * n);
        int tail_cnt = 0;
        long long last = 0;
        for (int j = n - 1; j >= 0; j--) {
            last += a[j];
            if (one_third == last) {
                tail[j] = ++tail_cnt;
            } else {
                tail[j] = tail_cnt;
            }
        }

        long long first = 0;
        for (int i = 0; i < n - 2; i++) {
            first += a[i];
            if (one_third == first) {
                result += tail[i + 2];
            }
        }

        free (tail);
    }

    printf ("%I64d\n", result);

    free (a);
    return 0;
}

