#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

void swap(char * x, char * y) {
    char tmp = *x;
    *x = *y;
    *y = tmp;
}

int main (int argc, char * argv[]) {
    int number_child, time;
    scanf ("%d %d\n", &number_child, &time);

    char * queue = malloc (sizeof(char) * number_child + 1);
    fgets(queue, number_child + 1, stdin);

    for (int t = 0; t < time; t++) {
        char last_i = queue[0];
        for (int i = 1; i < number_child && queue[i] != '\n'; i++) {
            if (last_i == 'B' && queue[i] == 'G') {
                last_i = queue[i];
                swap(&queue[i - 1], &queue[i]);
            } else {
                last_i = queue[i];
            }
        }
    }

    for (int i = 0; i < number_child && queue[i] != '\n'; i++) {
        printf ("%c", queue[i]);
    }
    printf ("\n");

    free (queue);
    return 0;
}

