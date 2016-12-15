#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, t;
    scanf ("%d %d", &n, &t);

    char * result = (char *) malloc (sizeof(char) * n);

    bool exist = true;
    if (n > 1) {
        if (t < 10) {
            for (int i = 0; i < n; i++) {
                result[i] = '0' + t;
            }
        } else {
            result[0] = '1';
            for (int i = 1; i < n; i++) {
                result[i] = '0';
            }
        }
    } else if (t < 10) {
        result[n - 1] = '0' + t;
    } else {
        printf ("-1");
        exist = false;
    }

    if (exist) {
        for (int i = 0; i < n; i++) {
            printf ("%c", result[i]);
        }
    }
    printf ("\n");

    free (result);
    return 0;
}

