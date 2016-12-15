#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, h, a, width = 0;
    scanf ("%d %d", &n, &h);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &a);
        if (a > h) {
            width += 2;
        } else {
            width += 1;
        }
    }

    printf ("%d\n", width);

    return 0;
}

