#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int number_errors, error;
    scanf ("%d", &number_errors);

    long long sum_all = 0L;
    for (int i = 0; i < number_errors; i++) {
        scanf ("%d", &error);
        sum_all += error;
    }
    long long sum_fixed_one = 0L;
    for (int i = 0; i < number_errors - 1; i++) {
        scanf ("%d", &error);
        sum_fixed_one += error;
    }
    long long sum_fixed_two = 0L;
    for (int i = 0; i < number_errors - 2; i++) {
        scanf ("%d", &error);
        sum_fixed_two += error;
    }
    printf ("%d\n", (int) (sum_all - sum_fixed_one));
    printf ("%d\n", (int) (sum_fixed_one - sum_fixed_two));

    return 0;
}

