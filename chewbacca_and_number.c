#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_DIGITS 19

int main (int argc, char * argv[]) {
    char * digits = (char *) malloc (sizeof(char) * (MAX_DIGITS + 2));
    fgets(digits, MAX_DIGITS + 2, stdin);

    int i = 0;
    if (digits[0] == '9') {
        i = 1;
    }

    for (; i < MAX_DIGITS && digits[i] != '\n'; i++) {
        int digit = digits[i] - '0';
        if (digit > 4) {
            digit = 9 - digit;
            digits[i] = digit + '0';
        }
    }
    for (int i = 0; i < MAX_DIGITS && digits[i] != '\n'; i++) {
        printf ("%c", digits[i]);
    }
    printf ("\n");

    free (digits);
    return 0;
}

