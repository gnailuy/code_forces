#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char * argv[]) {
    char input_char;
    do {
        scanf ("%c", &input_char);
        if ('H' == input_char || 'Q' == input_char || '9' == input_char) {
            printf ("YES\n");
            return 0;
        }
    } while ('\n' != input_char);

    printf ("NO\n");

    return 0;
}

