#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_NUMBER 1000000000000
#define MAX_PRIME 1000000

bool is_t_prime (long long num, long long * t_primes, int start, int end) {
    if (start > end) {
        return false;
    }

    int middle = (start + end) / 2;
    if (num == t_primes[middle]) {
        return true;
    } else if (num < t_primes[middle]) {
        return is_t_prime (num, t_primes, start, middle - 1);
    } else {
        return is_t_prime (num, t_primes, middle + 1, end);
    }
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int * primes = (int *) calloc (MAX_PRIME + 1, sizeof(int)); // Zero-initialized
    for (int i = 2; i <= MAX_PRIME; i++) {
        int max_j = MAX_PRIME / i;
        for (int j = 2; j <= max_j; j++) {
            primes[i * j] = 1; // primes[x] == 1 means primes[x] is not a prime
        }
    }

    int length = 0;
    long long * t_primes = (long long *) malloc (MAX_PRIME * sizeof(long long));
    for (int i = 2; i <= MAX_PRIME; i++) {
        if (0 == primes[i]) { // Is a prime
            t_primes[length++] = ((long long) i) * i;
        }
    }

    long long number;
    for (int i = 0; i < n; i++) {
        // scanf ("%I64d", &number);
        scanf ("%lld", &number);
        if (is_t_prime (number, t_primes, 0, length - 1)) {
            printf ("YES\n");
        } else {
            printf ("NO\n");
        }
    }

    free (primes);
    free (t_primes);
    return 0;
}

