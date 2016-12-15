#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_INPUT_LENGTH 100

char *hello = "hello\n";

int main (int argc, char * argv[]) {
    char in = ' ';
    int expected_pos = 0;

    for (int i = 0; i < MAX_INPUT_LENGTH && in != '\n'; i++) {
        scanf ("%c", &in);
        if (hello[expected_pos] == in) {
            expected_pos ++;
            if ('\n' == hello[expected_pos]) {
                printf ("YES\n");
                return 0;
            }
        }
    }

    printf ("NO\n");

    return 0;
}

