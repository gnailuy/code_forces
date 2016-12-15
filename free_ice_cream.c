#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, x, distress = 0;
    scanf ("%d %d\n", &n, &x);

    long long ice_creams = x;
    char plus_or_minus; int amount;
    for (int i = 0; i < n; i++) {
        scanf ("%c", &plus_or_minus);
        scanf ("%d\n", &amount);

        if ('+' == plus_or_minus) {
            ice_creams += amount;
        } else if ('-' == plus_or_minus) {
            if (ice_creams >= amount) {
                ice_creams -= amount;
            } else {
                distress++;
            }
        } else {
            printf ("Error!\n");
        }
    }

    // printf ("%I64d %d\n", ice_creams, distress);
    printf ("%lld %d\n", ice_creams, distress);

    return 0;
}
