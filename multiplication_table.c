#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, x, other, result = 0;
    scanf ("%d %d", &n, &x);

    for (int i = 1; i <= n && i * i <= x; i++) {
        if (x % i == 0) {
            other = x / i;
            if (other <= n) {
                if (other == i) result += 1;
                else result += 2;
            }
        }
    }

    printf ("%d\n", result);

    return 0;
}
