#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_WORD_LEN 1000
#define CASE_DIFF ('a' - 'A')

bool is_lower(char c) {
    if (c >= 'a' && c <= 'z') {
        return true;
    }
    return false;
}

char to_upper_case(char c) {
    if (is_lower(c)) {
        return c - CASE_DIFF;
    } else {
        return c;
    }
}

int main (int argc, char * argv[]) {
    char c;

    scanf ("%c", &c);
    printf ("%c", to_upper_case(c));

    for (int i = 1; i < MAX_WORD_LEN; i++) {
        scanf ("%c", &c);
        printf ("%c", c);
        if ('\n' == c) {
            break;
        }
    }

    return 0;
}

