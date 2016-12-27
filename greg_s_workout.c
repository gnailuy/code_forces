#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int max (int * array, int length) {
    int max = -1, index = -1;
    for (int i = 0; i < length; i++) {
        if (array[i] > max) {
            max = array[i];
            index = i;
        }
    }
    return index;
}

int main (int argc, char * argv[]) {
    int n, a;
    scanf ("%d", &n);

    int exercises[3] = { 0 };
    for (int i = 0; i < n; i++) {
        scanf ("%d", &a);
        exercises[i % 3] += a;
    }

    switch (max (exercises, 3)) {
        case 0:
            printf ("chest\n");
            break;
        case 1:
            printf ("biceps\n");
            break;
        case 2:
            printf ("back\n");
            break;
        default:
            printf ("ambiguous\n");
            break;
    }

    return 0;
}
