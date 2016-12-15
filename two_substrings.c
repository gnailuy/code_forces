#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LINE 100000

int main (int argc, char * argv[]) {
    char * line = (char *) malloc (sizeof(char) * MAX_LINE + 2);
    fgets(line, MAX_LINE + 2, stdin); // \n & \0

    int ab_count = 0, ba_count = 0;
    int * ab = (int *) malloc (sizeof(int) * (MAX_LINE - 1));
    int * ba = (int *) malloc (sizeof(int) * (MAX_LINE - 1));

    for (int i = 0; i < MAX_LINE && line[i] != '\n'; i++) {
        if ('A' == line[i] && 'B' == line[i + 1]) {
            ab[ab_count++] = i;
        }
        if ('B' == line[i] && 'A' == line[i + 1]) {
            ba[ba_count++] = i;
        }
    }

    bool has_abba = false;
    if (ab_count > 0 && ba_count > 0) {
        for (int i = 0; i < ab_count; i++) {
            for (int j = 0; j < ba_count; j++) {
                if (ab[i] + 1 != ba[j] && ba[j] + 1 != ab[i]) {
                    has_abba = true;
                    break;
                }
            }
            if (has_abba) {
                break;
            }
        }
    }

    if (has_abba) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (line);
    free (ab);
    free (ba);
    return 0;
}

