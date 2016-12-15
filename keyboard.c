#include <stdio.h>
#include <stdlib.h>

#define MAX_INPUT_LENGTH 100

const char * const layout[] = {
    "qwertyuiop",
    "asdfghjkl;",
    "zxcvbnm,./"
};
const int layout_length = 10;
const int layout_height = 3;

int main (int argc, char * argv[]) {
    char shifting;
    scanf ("%c\n", &shifting);

    int error;
    if ('L' == shifting) {
        error = 1;
    } else {
        error = -1;
    }

    char * input = (char *) malloc (sizeof(char) * MAX_INPUT_LENGTH);
    fgets(input, MAX_INPUT_LENGTH + 1, stdin);

    for (int i = 0; i < MAX_INPUT_LENGTH && input[i] != '\n'; i++) {
        for (int j = 0; j < layout_height; j++) {
            for (int k = 0; k < layout_length; k++) {
                if (layout[j][k] == input[i]) {
                    printf ("%c", layout[j][k + error]);
                }
            }
        }
    }
    printf ("\n");

    return 0;
}
