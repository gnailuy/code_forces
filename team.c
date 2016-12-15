#include <stdio.h>
#include <stdlib.h>

#define NUMBER_PLAYER 3 // Petya, Vasya and Tonya
#define MIN_SURE 2 // At least 2 of the players are sure about the solution

int main (int argc, char * argv[]) {
    int number_problems, sure_count;
    int number_implements = 0;
    int * views = malloc (sizeof(int) * NUMBER_PLAYER);

    scanf ("%d", &number_problems);

    for (int i = 0; i < number_problems; i++) {
        sure_count = 0;
        for (int j = 0; j < NUMBER_PLAYER; j++) {
            scanf ("%d", &views[j]);
            if (views[j] == 1) sure_count ++;
        }
        if (sure_count >= MIN_SURE) {
            number_implements ++;
        }
    }

    printf ("%d\n", number_implements);

    free (views);
    return 0;
}
