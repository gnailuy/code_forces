#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    char first, last, l;
    int lines, line_count = 0, letter_count = 0;
    char * word = (char *) malloc (sizeof(char) * (10 - 1));

    scanf ("%d", &lines);
    getchar ();

    while ( !feof(stdin) && line_count < lines ) {
        first = -1;
        last = -1;
        while ( scanf ("%c", &l) ) {
            if ( -1 == first ) {
                first = l;
                letter_count = 1;
            } else if ( '\n' != l && letter_count <= 10 ) {
                last = l;
                word[letter_count - 1] = l;
                letter_count++;
            } else if ( '\n' != l && letter_count > 10 ) {
                last = l;
                letter_count++;
            } else if ( letter_count > 10 ) {
                printf ("%c%d%c\n", first, letter_count - 2, last);
                break;
            } else {
                word[letter_count - 1] = '\0';
                printf ("%c%s\n", first, word);
                break;
            }
        }
        line_count ++;
    }

    free (word);
    return 0;
}
