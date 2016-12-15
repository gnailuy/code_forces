#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main (int argc, char * argv[]) {
    int n, t, next = 1;
    scanf ("%d %d", &n, &t);

    int * trans_code = (int *) malloc (sizeof(int) * n - 1);

    for (int i = 0; i < n - 1; i++) {
        scanf ("%d", &trans_code[i]);
    }

    bool answer = false;
    while (next < n) {
        next += trans_code[next - 1];
        if (t == next) {
            answer = true;
            break;
        } else if (t < next) {
            break;
        }
    }
    if (answer) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (trans_code);
    return 0;
}

