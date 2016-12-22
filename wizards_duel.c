#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int l, p, q;
    scanf ("%d", &l);
    scanf ("%d", &p);
    scanf ("%d", &q);

    printf ("%0.4f\n", ((double) l) * p / (p + q));

    return 0;
}

