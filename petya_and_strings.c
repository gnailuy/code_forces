#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE 100
#define DISTANCE 'a' - 'A'

int main (int argc, char * argv[]) {
    char * first = malloc (sizeof(char) * MAX_LINE + 1);
    char * second = malloc (sizeof(char) * MAX_LINE + 1);
    fgets(first, MAX_LINE + 2, stdin);
    fgets(second, MAX_LINE + 2, stdin);

    for (int i = 0; i < MAX_LINE && first[i] != '\n'; i++) {
        if (first[i] > 'Z') {
            first[i] -= DISTANCE;
        }
        if (second[i] > 'Z') {
            second[i] -= DISTANCE;
        }
        if (first[i] < second[i]) {
            printf ("-1\n");
            free (first);
            free (second);
            return 0;
        } else if (first[i] > second[i]) {
            printf ("1\n");
            free (first);
            free (second);
            return 0;
        }
    }

    printf ("0\n");
    free (first);
    free (second);
    return 0;
}

