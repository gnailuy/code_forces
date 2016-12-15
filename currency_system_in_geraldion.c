#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, note;
    scanf ("%d", &n);

    bool has_one = false;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &note);
        if (1 == note) {
            has_one = true;
            // Don't break, let's read all inputs
        }
    }

    if (has_one) {
        printf ("-1\n");
    } else {
        printf ("1\n");
    }

    return 0;
}

