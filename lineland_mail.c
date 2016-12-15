#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int max (int a, int b) {
    return a > b ? a : b;
}

int min (int a, int b) {
    return a < b ? a : b;
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int * cities = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &cities[i]);
    }

    printf ("%d %d\n",
            cities[1] - cities[0], cities[n - 1] - cities[0]);
    for (int i = 1; i < n - 1; i++) {
        printf ("%d %d\n",
                min (cities[i] - cities[i - 1], cities[i + 1] - cities[i]), 
                max (cities[i] - cities[0], cities[n - 1] - cities[i]));
    }
    printf ("%d %d\n",
            cities[n - 1] - cities[n - 2], cities[n - 1] - cities[0]);

    return 0;
}

