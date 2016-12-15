#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int days;
    scanf ("%d", &days);

    int * needs = (int *) malloc (sizeof(int) * days);
    int * price = (int *) malloc (sizeof(int) * days);
    for (int i = 0; i < days; i++) {
        scanf ("%d %d", &needs[i], &price[i]);
    }

    int money = 0;
    for (int i = 0; i < days; i++) {
        money += needs[i] * price[i];
        for (int j = i + 1; j < days; j++) {
            if (price[i] >= price[j]) {
                break;
            } else {
                money += needs[j] * price[i];
                needs[j] = 0;
            }
        }
    }

    printf ("%d\n", money);

    return 0;
}
