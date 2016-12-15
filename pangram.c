#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define FULL26 0x03FFFFFF

int letter_pos (char l) {
    if ('a' <= l && l <= 'z') {
        return l - 'a';
    } else if ('A' <= l && l <= 'Z') {
        return l - 'A';
    } else {
        return 26;
    }
}

int main (int argc, char * argv[]) {
    int length;
    scanf ("%d\n", &length);

    char * line = (char *) malloc (sizeof(char) * (length + 2));
    fgets(line, length + 2, stdin);

    if (length < 26) {
        printf ("NO\n");
        free (line);
        return 0;
    }

    int board = 0, pos;
    for (int i = 0; i < length; i++) {
        pos = 1 << letter_pos (line[i]);
        board |= pos;
    }

    if (board == FULL26) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (line);
    return 0;
}

