#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    printf ("I hate ");
    for (int i = 1; i < n; i++) {
        if (i % 2 == 0) {
            printf ("that I hate ");
        } else {
            printf ("that I love ");
        }
    }
    printf ("it\n");

    return 0;
}

