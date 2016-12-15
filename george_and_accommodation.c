#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char * argv[]) {
    int num_rooms, num_people, room_capacity;
    int result = 0;
    scanf ("%d\n", &num_rooms);

    for (int i = 0; i < num_rooms; i++) {
        scanf ("%d %d", &num_people, &room_capacity);
        if (room_capacity - num_people >= 2) {
            result += 1;
        }
    }

    printf ("%d\n", result);

    return 0;
}

