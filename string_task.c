#include <stdio.h>
#include <stdlib.h>

#define TRUE 1
#define FALSE 0

char to_lower_case (char c) {
    if ('A' <= c && 'Z' >= c) {
        return c - 'A' + 'a';
    } else {
        return c;
    }
}

int equal_ignore_case (char a, char b) {
    return to_lower_case (a) == to_lower_case (b);
}

int is_vowel (char c) {
    if (equal_ignore_case ('a', c) ||
        equal_ignore_case ('o', c) ||
        equal_ignore_case ('y', c) ||
        equal_ignore_case ('e', c) ||
        equal_ignore_case ('u', c) ||
        equal_ignore_case ('i', c)) {
        return TRUE;
    }
    return FALSE;
}

int main (int argc, char * argv[]) {
    char c;

    while ( scanf ("%c", &c) && '\n' != c ) {
        if (!is_vowel (c)) {
            printf (".%c", to_lower_case (c));
        }
    }
    printf ("\n");

    return 0;
}
