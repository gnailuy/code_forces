#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE 100
#define MAX_DIGIT 3

int main (int argc, char * argv[]) {
    int plus_sign_count = 0;
    char * line = malloc (sizeof(char) * MAX_LINE + 1);
    fgets(line, MAX_LINE + 1, stdin);

    int * summans = calloc (MAX_DIGIT, sizeof(int)); // Zero-initialized
    for (int i = 0; i < MAX_LINE && line[i] != '\n'; i++) {
        if (line[i] != '+') {
            summans[line[i] - '1'] += 1;
        } else {
            plus_sign_count ++;
        }
    }
    for (int i = 0; i < MAX_DIGIT; i++) {
        for (int j = 0; j < summans[i]; j++) {
            printf ("%c", '1' + i);
            if (plus_sign_count-- > 0) {
                printf ("+");
            }
        }
    }

    printf ("\n");

    free (line);
    free (summans);
    return 0;
}

