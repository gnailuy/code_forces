#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char * argv[]) {
    int price_first, money, num_bananas;
    long result = 0;
    scanf ("%d %d %d", &price_first, &money, &num_bananas);

    result = (1 + num_bananas) * num_bananas * price_first / 2 - money;
    if (result < 0) {
        result = 0;
    }

    printf ("%d\n", result);

    return 0;
}

