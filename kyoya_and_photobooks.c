#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LENGTH 21

int main (int argc, char * argv[]) {
    char in, last_in = ' ';
    int length = 0;
    for (int i = 0; i < MAX_LENGTH; i++) {
        scanf ("%c", &in);
        if ('\n' == in) break;

        length++;
    }

    printf ("%d\n", (length + 1) * 26 - length);

    return 0;
}

