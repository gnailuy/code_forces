#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, edge, turn, sum_inner_angle;
    bool ok;

    scanf ("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &turn);
        edge = 2;
        ok = false;
        do {
            edge++; // Start from triangle
            sum_inner_angle = (edge - 2) * 180;
            if (sum_inner_angle == turn * edge) {
                ok = true;
                break;
            }
        } while (sum_inner_angle <= turn * edge);
        if (ok) {
            printf ("YES\n");
        } else {
            printf ("NO\n");
        }
    }

    return 0;
}

