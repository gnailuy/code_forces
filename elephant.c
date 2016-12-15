#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_STEP 5

int main (int argc, char * argv[]) {
    int target, steps = 0;
    scanf ("%d", &target);

    for (int i = MAX_STEP; i > 0; i--) {
        if (target >= i) {
            steps += target / i;
            target = target % i;
        }
        if (0 == target) break;
    }

    printf ("%d\n", steps);

    return 0;
}

