#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_WORD_LEN 100

int main (int argc, char * argv[]) {
    char * berlandish_word = (char *) malloc (sizeof(char) * (MAX_WORD_LEN + 2));
    char * birlandish_word = (char *) malloc (sizeof(char) * (MAX_WORD_LEN + 2));
    fgets(berlandish_word, MAX_WORD_LEN + 2, stdin);
    fgets(birlandish_word, MAX_WORD_LEN + 2, stdin);

    int ber_len = strlen (berlandish_word) - 1;
    int bir_len = strlen (birlandish_word) - 1;

    if (ber_len != bir_len) {
        printf ("NO\n");
        free (berlandish_word);
        free (birlandish_word);
        return 0;
    }

    for (int i = 0; i < ber_len; i++) {
        if (berlandish_word [i] != birlandish_word[ber_len - i - 1]) {
            printf ("NO\n");
            free (berlandish_word);
            free (birlandish_word);
            return 0;
        }
    }

    printf ("YES\n");
    free (berlandish_word);
    free (birlandish_word);
    return 0;
}

