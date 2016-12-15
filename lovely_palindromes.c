#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_LEN 100001

int main (int argc, char * argv[]) {
    char * num = (char *) malloc (sizeof(char) * (MAX_LEN + 2));
    fgets (num, MAX_LEN + 2, stdin);

    int i;
    for (i = 0; i < MAX_LEN && num[i] != '\n'; i++) {
        printf ("%c", num[i]);
    }
    for (i = i - 1; i >= 0; i--) {
        printf ("%c", num[i]);
    }

    printf ("\n");
}

