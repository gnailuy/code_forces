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

char to_other_case(char c) {
    if (is_upper(c)) {
        return c + CASE_DIFF;
    } else if (is_lower(c)){
        return c - CASE_DIFF;
    } else {
        return c;
    }
}

int main (int argc, char * argv[]) {
    char * word = (char *) malloc (sizeof(char) * (MAX_WORD_LEN + 1));
    fgets(word, MAX_WORD_LEN + 1, stdin);

    bool is_accident = true;
    for (int i = 1; i < MAX_WORD_LEN && word[i] != '\n'; i++) {
        if (is_lower(word[i])) {
            is_accident = false;
            break;
        }
    }

    if (is_accident) {
        for (int i = 0; i < MAX_WORD_LEN && word[i] != '\n'; i++) {
            printf ("%c", to_other_case(word[i]));
        }
    } else {
        for (int i = 0; i < MAX_WORD_LEN && word[i] != '\n'; i++) {
            printf ("%c", word[i]);
        }
    }
    printf ("\n");

    free (word);
    return 0;
}

