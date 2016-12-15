#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    int max, min, current, amazing = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &current);
        if (i == 0) {
            max = min = current;
        } else {
            if (current > max) {
                max = current;
                amazing ++;
            } else if (current < min) {
                min = current;
                amazing ++;
            }
        }
    }

    printf ("%d\n", amazing);

    return 0;
}

