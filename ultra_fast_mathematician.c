#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LENGTH 100

int main (int argc, char * argv[]) {
    char * f = (char *) malloc (sizeof(char) * MAX_LENGTH + 2);
    char * s = (char *) malloc (sizeof(char) * MAX_LENGTH + 2);

    fgets(f, MAX_LENGTH + 2, stdin);
    fgets(s, MAX_LENGTH + 2, stdin);

    for (int i = 0; i < MAX_LENGTH && f[i] != '\n'; i++) {
        if (f[i] != s[i]) {
            printf ("1");
        } else {
            printf ("0");
        }
    }

    printf ("\n");

    free (f);
    free (s);
    return 0;
}

