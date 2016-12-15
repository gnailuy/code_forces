#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define NUM_PLAYER 5

int main (int argc, char * argv[]) {
    int amount, sum = 0;
    for (int i = 0; i < NUM_PLAYER; i++) {
        scanf ("%d", &amount);
        sum += amount;
    }

    if (sum % NUM_PLAYER == 0 && sum > 0) {
        printf ("%d\n", sum / NUM_PLAYER);
    } else {
        printf ("-1\n");
    }

    return 0;
}

