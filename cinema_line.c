#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, bill;
    scanf ("%d", &n);

    int num_25 = 0, num_50 = 0;

    bool smooth = true;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &bill);
        if (25 == bill) {
            num_25++;
        } else if (50 == bill) {
            if  (num_25 > 0) {
                num_25--;
                num_50++;
            } else {
                smooth = false;
                break;
            }
        } else if (100 == bill) {
            if (num_50 > 0 && num_25 > 0) {
                num_50--;
                num_25--;
            } else if (num_25 > 2) {
                num_25 -= 3;
            } else {
                smooth = false;
                break;
            }
        }
    }

    if (smooth) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    return 0;
}

