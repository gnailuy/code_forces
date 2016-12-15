#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define NUMBER_OF_LUCKY_NUMBERS 13

int lucky_numbers[NUMBER_OF_LUCKY_NUMBERS]
        = {4, 7,
           44, 47, 74,
           444, 447, 474, 477, 744, 747, 774, 777};

int main (int argc, char * argv[]) {
    int number;
    scanf ("%d", &number);

    for (int i = 0; i < NUMBER_OF_LUCKY_NUMBERS; i++) {
        if (0 == number % lucky_numbers[i]) {
            printf ("YES\n");
            return 0;
        }
    }
    printf ("NO\n");

    return 0;
}

