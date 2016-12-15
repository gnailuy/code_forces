#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int a, b, result;
    scanf ("%d %d", &a, &b);

    int r = 0, m;
    while (true) {
        result += a; // Burn all candles
        m = a + r; // Materials for new candle
        if (m < b) {
            break;
        }

        a = m / b; // Make new candles out of m
        r = m % b; // And remain some r
    }

    printf ("%d\n", result);

    return 0;
}

