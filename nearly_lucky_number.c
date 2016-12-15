#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_LINE 19

int main (int argc, char * argv[]) {
    int number_lucky_digits = 0;
    char * line = malloc (sizeof(char) * MAX_LINE + 1);
    fgets(line, MAX_LINE + 1, stdin);

    for (int i = 0; i < MAX_LINE && line[i] != '\n'; i++) {
        if (line[i] == '4' || line[i] == '7') {
            number_lucky_digits ++;
        }
    }
    if (number_lucky_digits == 4 || number_lucky_digits == 7) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (line);
    return 0;
}

