#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, c;
    scanf ("%d %d", &n, &c);

    int last_price, current_price, max_diff = 0;
    scanf ("%d", &last_price);
    for (int i = 1; i < n; i++) {
        scanf ("%d", &current_price);
        int diff = last_price - current_price;
        if (diff >= 0 && diff > max_diff) {
            max_diff = diff;
        }

        last_price = current_price;
    }

    printf ("%d\n", max_diff >= c ? max_diff - c : 0);

    return 0;
}
