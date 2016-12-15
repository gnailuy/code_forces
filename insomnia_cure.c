#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define HIT_TYPE 4

int main (int argc, char * argv[]) {
    int hit[HIT_TYPE];
    int d;
    int result = 0;

    for (int i = 0; i < HIT_TYPE; i++) {
        scanf ("%d", &hit[i]);
    }
    scanf ("%d", &d);

    for (int i = 1; i <= d; i++) {
        for (int j = 0; j < HIT_TYPE; j++) {
            if (i % hit[j] == 0) {
                result += 1;
                break;
            }
        }
    }

    printf ("%d\n", result);

    return 0;
}

