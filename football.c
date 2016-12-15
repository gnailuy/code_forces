#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DANGEROUS_THRESHOLD 7
#define MAX_LINE 100

int main (int argc, char * argv[]) {
    int continuous_count = 0;
    char current_team = 'n';

    char * seq = malloc (sizeof(char) * MAX_LINE + 1);
    fgets(seq, MAX_LINE + 1, stdin);

    for (int i = 0; i < MAX_LINE && seq[i] != '\n'; i++) {
        if (seq[i] != current_team) {
            continuous_count = 1;
        } else {
            continuous_count += 1;
            if (continuous_count >= DANGEROUS_THRESHOLD) {
                printf ("YES\n");

                free (seq);
                return 0;
            }
        }
        current_team = seq[i];
    }

    printf ("NO\n");

    free (seq);
    return 0;
}

