#include <stdio.h>
#include <stdlib.h>

#define C_STAT 3

int main (int argc, char * argv[]) {
    int number_statments;
    int X = 0;
    char statment[C_STAT + 1]; // For the tailing '\0' in an input string

    scanf ("%d", &number_statments);

    for (int i = 0; i < number_statments; i++) {
        scanf ("%s", statment);
        if (statment[1] == '-') {
            --X;
        } else if (statment[1] == '+') {
            ++X;
        }
    }

    printf ("%d\n", X);

    return 0;
}

