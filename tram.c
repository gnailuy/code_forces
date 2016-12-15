#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int num_stations;
    int current = 0;
    int capacity = 0;
    int enter = 0, exit = 0;

    scanf ("%d", &num_stations);

    for (int i = 0; i < num_stations; i++) {
        scanf ("%d %d", &exit, &enter);
        current -= exit;
        current += enter;
        if (current > capacity) {
            capacity = current;
        }
    }

    printf ("%d\n", capacity);

    return 0;
}

