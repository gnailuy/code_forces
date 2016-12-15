#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, v;
    scanf ("%d %d", &n, &v);

    int * sellers = (int *) calloc (n, sizeof(int)); // Zero-initialized

    int items, price, cnt = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &items);
        for (int j = 0; j < items; j++) {
            scanf ("%d", &price);
            if (price < v && sellers[i] == 0) {
                sellers[i] = 1;
                cnt++;
            }
        }
    }

    printf ("%d\n", cnt);
    if (cnt > 0) {
        for (int i = 0; i < n; i++) {
            if (1 == sellers[i]) {
                printf ("%d ", i + 1);
            }
        }
        printf ("\n");
    }

    free (sellers);
    return 0;
}

