#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char * argv[]) {
    int number_of_stones = 0;
    int take_out = 0;

    scanf ("%d\n", &number_of_stones);
    char * seq = malloc (sizeof(char) * number_of_stones + 1);
    fgets(seq, number_of_stones + 1, stdin);

    char current_color = seq[0];
    for (int i = 1; i < number_of_stones && seq[i] != '\n'; i++) {
        if (current_color != seq[i]) {
            current_color = seq[i];
            continue;
        } else {
            take_out ++;
        }
    }

    printf ("%d\n", take_out);

    free (seq);
    return 0;
}

