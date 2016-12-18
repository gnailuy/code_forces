#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    int team = 0;
    while (true) {
        if (n >= m) {
            if (n > 1 && m > 0) {
                team++;
                n -= 2;
                m -= 1;
            } else {
                break;
            }
        } else {
            if (m > 1 && n > 0) {
                team++;
                m -= 2;
                n -= 1;
            } else {
                break;
            }
        }
    }

    printf ("%d\n", team);

    return 0;
}

