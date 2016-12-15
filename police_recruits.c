#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, event, polices = 0, untreated = 0;
    scanf ("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &event);
        if (event > 0) {
            polices += event;
        } else if (-1 == event) {
            if (polices > 0) {
                polices--;
            } else {
                untreated++;
            }
        }
    }

    printf ("%d\n", untreated);

    return 0;
}

