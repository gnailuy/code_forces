#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void print_line_i (int n, int i) {
    for (int j = 0; j < n - i; j++) {
        printf ("  ");
    }
    for (int j = 0; j < i; j++) {
        printf ("%d ", j);
    }
    for (int j = 0; j <= i; j++) {
        if (j == i) {
            printf ("%d", i - j);
        } else {
            printf ("%d ", i - j);
        }
    }
    printf ("\n");
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    for (int i = 0; i <= n; i++) {
        print_line_i (n, i);
    }
    
    for (int i = n - 1; i >= 0; i--) {
        print_line_i (n, i);
    }

    return 0;
}

