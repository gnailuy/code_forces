#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, target, sum = 0;
    scanf ("%d", &n); // Guaranteed even

    int * numbers = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &numbers[i]);
        sum += numbers[i];
    }
    target = sum / (n / 2);

    for (int i = 0; i < n; i++) {
        if (numbers[i] != -1) {
            for (int j = i + 1; j < n; j++) {
                if (numbers[i] + numbers[j] == target) {
                    printf ("%d %d\n", i + 1, j + 1);
                    numbers[i] = numbers[j] = -1;
                }
            }
        }
    }

    free (numbers);
    return 0;
}

