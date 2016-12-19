#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int a, b, c;
    scanf ("%d %d %d", &a, &b, &c);

    bool result;
    if (c == 0) {
        if (a == b) {
            result = true;
        } else {
            result = false;
        }
    } else if (c > 0) {
        if (a > b) {
            result = false;
        } else {
            result = (b - a) % c == 0;
        }
    } else if (c < 0) {
        if (a < b) {
            result = false;
        } else {
            result = (a - b) % (-c) == 0;
        }
    }

    if (result) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    return 0;
}
