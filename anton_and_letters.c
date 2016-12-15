#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LINE 1000

int main (int argc, char * argv[]) {
    int index, result = 0;
    char * line = (char *) malloc (sizeof(char) * MAX_LINE + 2);
    fgets(line, MAX_LINE + 2, stdin);

    unsigned int mask, filter = 0;

    for (int i = 0; i < MAX_LINE && line[i] != '\n'; i++) {
        index = line[i] - 'a';
        if (index >= 0 && index <= 25) { // Small letters
            mask = 1 << index;
            if ((filter & mask) == 0) {
                result ++;
                filter += mask;
            }
        }
    }

    printf ("%d\n", result);

    free (line);
    return 0;
}

