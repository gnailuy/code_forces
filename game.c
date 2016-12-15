#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n1, n2, k1, k2;
    scanf ("%d %d %d %d", &n1, &n2, &k1, &k2);

    if (n1 > n2) {
        printf ("First\n");
    } else {
        printf ("Second\n");
    }

    return 0;
}

