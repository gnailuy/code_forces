#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int k, r, burles, shovels = 0;

    scanf ("%d %d", &k, &r);

    do {
        shovels++;
        burles = shovels * k;
    } while (burles % 10 != 0 && burles % 10 != r);

    printf ("%d\n", shovels);

    return 0;
}

