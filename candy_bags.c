#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, half;
    scanf ("%d", &n); // n is even
    half = n / 2;

    int index = 1, xedni = n * n;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < half; j++) {
            printf ("%d ", index++);
        }
        for (int j = 0; j < half; j++) {
            printf ("%d ", xedni--);
        }
        printf ("\n");
    }

    return 0;
}

