#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_WORD_LEN 200

int main (int argc, char * argv[]) {
    char * seq = (char *) malloc (sizeof(char) * (MAX_WORD_LEN + 1));
    fgets(seq, MAX_WORD_LEN + 1, stdin);

    bool new_word = false;
    for (int i = 0; i < MAX_WORD_LEN && seq[i] != '\n'; i++) {
        if ('W' == seq[i] && 'U' == seq[i + 1] && 'B' == seq[i + 2]) {
            if (new_word && '\n' != seq[i + 3]) {
                new_word = false;
                printf (" ");
            }
            i += 2; // Then i++
        } else {
            new_word = true;
            printf ("%c", seq[i]);
        }
    }
    printf ("\n");

    free (seq);
    return 0;
}

