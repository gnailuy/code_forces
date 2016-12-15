#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_YEAR_DIGITS 5

bool is_beautiful_year(int year) {
    int * digits = (int *) malloc (sizeof(int) * MAX_YEAR_DIGITS);
    int idx = 0;
    int digit;

    while (year > 0) {
        digit = year % 10;
        for (int i = 0; i < idx; i++) {
            if (digit == digits[i]) {
                free (digits);
                return false;
            }
        }

        digits[idx++] = digit;
        year /= 10;
    }

    free (digits);
    return true;
}

int main (int argc, char * argv[]) {
    int year;
    scanf ("%d", &year);

    while (true) {
        year += 1;
        if (is_beautiful_year(year)) {
            printf ("%d\n", year);
            return 0;
        }
    }

    return 0;
}

