#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int base_holidays = (n / 7) * 2;
    int remain = n % 7;
    int min_holidays = base_holidays, max_holidays = base_holidays;
    if (1 == remain) {
        max_holidays += 1;
    } else if (6 == remain) {
        min_holidays += 1;
        max_holidays += 2;
    } else if (0 != remain) {
        max_holidays += 2;
    }

    printf ("%d %d\n", min_holidays, max_holidays);

    return 0;
}
