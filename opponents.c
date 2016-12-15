#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, d;
    scanf ("%d %d\n", &n, &d);

    char opp;
    int max_count = 0, current_count = 0;
    for (int i = 0; i < d; i++) {
        bool will_win = false;
        for (int j = 0; j <= n; j++) { // '\n'
            scanf ("%c", &opp);
            if ('0' == opp) {
                will_win = true;
            }
        }
        if (will_win) {
            current_count++;
        } else {
            if (current_count > max_count) {
                max_count = current_count;
            }
            current_count = 0;
        }
    }
    if (current_count > max_count) {
        max_count = current_count;
    }

    printf ("%d\n", max_count);

    return 0;
}

