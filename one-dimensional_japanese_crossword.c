#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    char * line = (char *) malloc (sizeof(char) * (n + 2));
    fgets(line, n + 2, stdin);

    int block_size, block_count = 0;
    int * blocks = (int *) malloc (sizeof(int) * (n / 2 + 1));

    bool in_block = false;
    for (int i = 0; i < n && line[i] != '\n'; i++) {
        if (line[i] == 'B') {
            if (in_block) {
                block_size += 1;
            } else {
                block_size = 1;
                in_block = true;
            }
        } else {
            if (in_block) {
                blocks[block_count++] = block_size;
                in_block = false;
            }
        }
    }
    if (in_block) {
        blocks[block_count++] = block_size;
    }

    printf ("%d\n", block_count);
    for (int i = 0; i < block_count; i++) {
        printf ("%d ", blocks[i]);
    }
    if (block_count > 0) printf ("\n");

    free (line);
    free (blocks);
    return 0;
}

