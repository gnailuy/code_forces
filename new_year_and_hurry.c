#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define TIME 240

int main (int argc, char * argv[]) {
    int n, k;
    scanf ("%d %d", &n, &k);

    int remain = TIME - k, result = 0;
    for (int i = 1; i <= n; i++) {
        int need = 5 * i;
        if (remain >= need) {
            remain -= need;
            result++;
        } else {
            break;
        }
    }

    printf ("%d\n", result);

    return 0;
}

