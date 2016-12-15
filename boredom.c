#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int * numbers = malloc (sizeof(int) * n);
    int max_number = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &numbers[i]);
        if (numbers[i] > max_number) {
            max_number = numbers[i];
        }
    }

    int * counts = calloc (max_number + 1, sizeof(int)); // Zero-initialized
    for (int i = 0; i < n; i++) {
        counts[numbers[i]] ++;
    }
    free (numbers);

    long long last_max[2] = { 0, counts[1] };
    long long result = 0;
    for (int i = 2; i <= max_number; i++) {
        result = last_max[0] + counts[i] * (long long) i;
        if (last_max[1] > result) {
            result = last_max[1];
        }
        last_max[0] = last_max[1];
        last_max[1] = result;
    }
    printf ("%I64d\n", result);

    free (counts);
    return 0;
}

