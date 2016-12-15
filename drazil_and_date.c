#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int abs (int x) {
    return x > 0 ? x : -x;
}

int main (int argc, char * argv[]) {
    int a, b, s;
    scanf ("%d %d %d", &a, &b, &s);

    if (abs(a) + abs(b) <= s && (s - abs(a) - abs(b)) % 2 == 0) {
        printf ("Yes\n");
    } else {
        printf ("No\n");
    }

    return 0;
}

