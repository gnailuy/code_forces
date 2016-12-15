#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void swap(int *a, int *b) {
    int c = *a;
    *a = *b;
    *b = c;
}

int main (int argc, char * argv[]) {
    int number_of_coins = 0;
    int take_out = 0;
    int total_money = 0;

    scanf ("%d\n", &number_of_coins);
    int * seq = malloc (sizeof(int) * number_of_coins + 1);
    for (int i = 0; i < number_of_coins; i++) {
        scanf ("%d", &seq[i]);
        total_money += seq[i];
    }
    for (int i = 0; i < number_of_coins; i++) {
        for (int j = i + 1; j < number_of_coins; j++) {
            if (seq[i] < seq[j]) {
                swap (&seq[i], &seq[j]);
            }
        }
    }

    int current_money = 0;
    double half_money = total_money / 2.0L;
    for (int i = 0; i < number_of_coins; i++) {
        current_money += seq[i];
        take_out += 1;
        if (current_money > half_money) {
            break;
        }
    }
    printf ("%d\n", take_out);

    free (seq);
    return 0;
}

