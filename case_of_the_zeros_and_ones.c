#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, zero_count = 0, one_count = 0;
    scanf ("%d\n", &n);

    char input;
    do {
        scanf ("%c", &input);
        if ('0' == input) {
            zero_count += 1;
        } else if ('1' == input) {
            one_count += 1;
        }
    } while (input != '\n');

    printf ("%d\n", zero_count > one_count ? zero_count - one_count
                                           : one_count - zero_count);

    return 0;
}

