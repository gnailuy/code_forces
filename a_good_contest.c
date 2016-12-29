#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define RED 2400

int main (int argc, char * argv[]) {
    char trash_bin;
    int n, before, after;
    scanf ("%d\n", &n);

    bool good = false;
    for (int i = 0; i < n; i++) {
        while (true) { // Ignore the name
            scanf ("%c", &trash_bin);
            if (' ' == trash_bin) break;
        }
        scanf ("%d %d", &before, &after);
        if (before >= RED && after > before) {
            good = true;
        }
    }
    if (good) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    return 0;
}

