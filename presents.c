#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int number_friends;
    scanf ("%d", &number_friends);

    int * seq = malloc (sizeof(int) * number_friends);
    for (int i = 0; i < number_friends; i++) {
        scanf ("%d", &seq[i]);
    }

    int * res = malloc (sizeof(int) * number_friends);
    for (int i = 0; i < number_friends; i++) {
        res[seq[i] - 1] = i + 1;
    }

    for (int i = 0; i < number_friends; i++) {
        printf ("%d ", res[i]);
    }
    printf ("\n");

    free (seq);
    free (res);
    return 0;
}

