#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_INPUT_LENGTH 100

int main (int argc, char * argv[]) {
    int number_distinct_char = 0;

    char * name = malloc (sizeof(char) * MAX_INPUT_LENGTH + 1);
    fgets(name, MAX_INPUT_LENGTH + 1, stdin);

    for (int i = 0; i < MAX_INPUT_LENGTH && name[i] != '\n'; i++) {
        bool has_dup = false;
        for (int j = i + 1; j < MAX_INPUT_LENGTH && name[j] != '\n'; j++) {
            if (name[i] == name[j]) {
                has_dup = true;
                break;
            }
        }
        if (!has_dup) {
            number_distinct_char += 1;
        }
    }

    if (0 == number_distinct_char % 2) {
        printf ("CHAT WITH HER!\n");
    } else {
        printf ("IGNORE HIM!\n");
    }

    free (name);
    return 0;
}

