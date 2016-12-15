#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int abs (int a) {
    return a > 0 ? a : -a;
}

int main (int argc, char * argv[]) {
    int x1, y1, x2, y2, edge_length;
    scanf ("%d %d %d %d", &x1, &y1, &x2, &y2);

    if (x1 == x2) {
        edge_length = abs (y1 - y2);
        printf ("%d %d %d %d\n",
                x1 + edge_length, y1, x2 + edge_length, y2); // x3, y3, x4, y4
    } else if (y1 == y2) {
        edge_length = abs (x1 - x2);
        printf ("%d %d %d %d\n",
                x1, y1 + edge_length, x2, y2 + edge_length); // x3, y3, x4, y4
    } else if (abs (x1 - x2) == abs (y1 - y2)) {
        printf ("%d %d %d %d\n", x1, y2, x2, y1); // x3, y3, x4, y4
    } else {
        printf ("-1\n");
    }

    return 0;
}

