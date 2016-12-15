#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_WORD_LEN 100
#define CASE_DIFF ('a' - 'A')

bool is_lower(char c) {
    if (c >= 'a' && c <= 'z') {
        return true;
    }
    return false;
}

bool is_upper(char c) {
    if (c >= 'A' && c <= 'Z') {
        return true;
    }
    return false;
}

char to_lower(char c) {
    if (is_upper(c)) {
        return c + CASE_DIFF;
    } else {
        return c;
    }
}

char to_upper(char c) {
    if (is_lower(c)) {
        return c - CASE_DIFF;
    } else {
        return c;
    }
}

int main (int argc, char * argv[]) {
    char * word = (char *) malloc (sizeof(char) * (MAX_WORD_LEN + 2));
    fgets(word, MAX_WORD_LEN + 2, stdin);

    int upper_count = 0, length = 0;
    for (int i = 0; i < MAX_WORD_LEN && word[i] != '\n'; i++) {
        if (is_upper(word[i])) {
            upper_count += 1;
        }
        length += 1;
    }

    if (2 * upper_count <= length) {
        for (int i = 0; i < MAX_WORD_LEN && word[i] != '\n'; i++) {
            printf ("%c", to_lower(word[i]));
        }
    } else {
        for (int i = 0; i < MAX_WORD_LEN && word[i] != '\n'; i++) {
            printf ("%c", to_upper(word[i]));
        }
    }
    printf ("\n");

    free (word);
    return 0;
}

