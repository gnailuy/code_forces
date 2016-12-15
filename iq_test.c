#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, number, result;
    scanf ("%d", &n);

    int even = 0, odd = 0;
    int even_index = 0, odd_index = 0;
    for (int i = 0; i < n; i++) {
        scanf ("%d", &number);
        if ((even > 0 && odd > 1)
            || (even > 1 && odd > 0)) {
            continue;
        }
        if (number % 2 == 0) {
            even += 1;
            even_index = i + 1;
        } else {
            odd += 1;
            odd_index = i + 1;
        }
    }
    if (even > 1) {
        printf ("%d\n", odd_index);
    } else {
        printf ("%d\n", even_index);
    }

    return 0;
}

