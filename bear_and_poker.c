#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int find_factor (int n) { // Find factors other than 2 and 3
    while (n % 2 == 0 || n % 3 == 0) {
        if (n % 2 == 0) {
            n /= 2;
        }
        if (n % 3 == 0) {
            n /= 3;
        }
    }
    return n;
}

int main (int argc, char * argv[]) {
    int n, a;
    scanf ("%d", &n);

    int factor = -1;
    bool ok = true;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &a);
        if (ok) {
            int f = find_factor (a);
            if (factor == -1) {
                factor = f;
            } else if (factor != f) {
                ok = false;
            }
        }
    }

    if (ok) {
        printf ("Yes\n");
    } else {
        printf ("No\n");
    }

    return 0;
}

