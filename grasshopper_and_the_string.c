#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_STR_LENGTH 100

int main (int argc, char * argv[]) {
    char in = ' ';
    int max_jump = 0, jump = 0, last_pos = 0;

    for (int i = 1; i <= MAX_STR_LENGTH + 1; i++) {
        scanf ("%c", &in);
        if ('A' == in || 'E' == in ||
                'I' == in || 'O' == in ||
                'U' == in || 'Y' == in || '\n' == in) {
            jump = i - last_pos;
            if (max_jump < jump) {
                max_jump = jump;
            }
            last_pos = i;
        }
        if ('\n' == in) {
            break;
        }
    }

    printf ("%d\n", max_jump);

    return 0;
}

