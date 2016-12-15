#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int number_days, max_non_decreasing_len = 1;
    scanf ("%d", &number_days);
    if (1 == number_days) { // 1 <= number_days <= 10^5
        printf ("%d\n", max_non_decreasing_len);
        return 0;
    }

    int * seq = malloc (sizeof(int) * number_days);
    for (int i = 0; i < number_days; i++) {
        scanf ("%d", &seq[i]);
    }

    int current_non_decreasing_len = 1;
    for (int i = 1; i < number_days; i++) {
        if (seq[i] >= seq[i - 1]) {
            current_non_decreasing_len += 1;
        } else {
            if (current_non_decreasing_len > max_non_decreasing_len) {
                max_non_decreasing_len = current_non_decreasing_len;
            }
            current_non_decreasing_len = 1;
        }
    }
    if (current_non_decreasing_len > max_non_decreasing_len) {
        max_non_decreasing_len = current_non_decreasing_len;
    }

    printf ("%d\n", max_non_decreasing_len);

    free (seq);
    return 0;
}

