#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int result = 0, need = 1;

    while (n >= need) {
        result ++;
        n -= need;
        need += (result + 1);
    }

    printf ("%d\n", result);

    return 0;
}

