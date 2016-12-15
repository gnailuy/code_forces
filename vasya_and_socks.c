#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    int result = 0, left_days = 0, buy_days;
    while (true) {
        result += n;
        buy_days = n + left_days;
        if (buy_days < m) {
            break;
        }
        n = buy_days / m;
        left_days = buy_days % m;
    }
    printf ("%d\n", result);

    return 0;
}

