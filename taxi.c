#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_PASSENGERS 4

int min_group(int *groups) {
    for (int i = 0; i < MAX_PASSENGERS; i++) {
        if (groups[i] > 0) {
            return i + 1;
        }
    }
    return 0; // Invalid
}

int main(int argc, char * argv[]) {
    long group_number;
    int s;
    int * groups;

    scanf("%ld", &group_number);
    groups = (int *) malloc(sizeof(int) * MAX_PASSENGERS);
    memset(groups, 0, sizeof(int) * MAX_PASSENGERS);
    for (int i = 0; i < group_number; i++) {
        scanf("%d", &s);
        groups[s - 1]++;
    }

    int car_count = 0;
    int on_board_count = 0;
    int remain_seats = 0;

    while (on_board_count < group_number) {
        if (remain_seats >= min_group(groups)) {
            for (int i = remain_seats - 1; i >= 0; i--) {
                if (groups[i] > 0) {
                    remain_seats -= (i + 1);
                    groups[i]--;
                    on_board_count++;
                    break;
                }
            }
        } else {
            car_count++;
            remain_seats = 4;
        }
    }

    printf("%d\n", car_count);

    free (groups);
    return 0;
}

