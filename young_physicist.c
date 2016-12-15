#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int number_forces, fx, fy, fz;
    int x = 0, y = 0, z = 0;
    scanf ("%d", &number_forces);

    for (int i = 0; i < number_forces; i++) {
        scanf ("%d %d %d", &fx, &fy, &fz);
        x += fx;
        y += fy;
        z += fz;
    }

    if (x != 0 || y != 0 || z != 0) {
        printf ("NO\n");
    } else {
        printf ("YES\n");
    }

    return 0;
}

