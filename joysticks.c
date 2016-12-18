#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int a, b;
    scanf ("%d %d", &a, &b);

    int time = 0;
    while (true) {
        if (a >= b) {
            if (a >= 2 && b > 0) {
                time++;
                a -= 2;
                b += 1;
            } else {
                break;
            }
        } else {
            if (b >= 2 && a > 0) {
                time++;
                b -= 2;
                a += 1;
            } else {
                break;
            }
        }
    }

    printf ("%d\n", time);

    return 0;
}

