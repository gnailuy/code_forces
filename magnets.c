#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define POLES 2

int main (int argc, char * argv[]) {
    int n, groups = 0;
    char last_pole = 'x';
    scanf ("%d\n", &n);

    char in[POLES + 2]; // Add 1 for the tailing '\n'
    for (int i = 0; i < n; i++) {
        fgets(in, POLES + 2, stdin);
        if (last_pole == in[0] || last_pole == 'x') {
            groups += 1;
        }
        last_pole = in[1];
    }

    printf ("%d\n", groups);

    return 0;
}

