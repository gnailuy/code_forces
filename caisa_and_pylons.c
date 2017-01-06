#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int h, max_h = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &h);
        if (h > max_h) {
            max_h = h;
        }
    }

    printf ("%d\n", max_h);

    return 0;
}
