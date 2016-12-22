#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_LEN 100000

int main (int argc, char * argv[]) {
    int a[4];
    scanf ("%d %d %d %d\n", &a[0], &a[1], &a[2], &a[3]);

    char * process = (char *) malloc (sizeof(char) * (MAX_LEN + 2));
    fgets(process, MAX_LEN + 2, stdin);

    int calories = 0;
    for (int i = 0; i < MAX_LEN && process[i] != '\n'; i++) {
        calories += a[process[i] - '1'];
    }

    printf ("%d\n", calories);

    free (process);
    return 0;
}

