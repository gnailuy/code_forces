#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define DAYS_IN_2016 366
#define WEEKDAY_20160101 5

int main (int argc, char * argv[]) {
    int x;
    scanf ("%d", &x);

    bool week;
    int index = 0; char c;
    do {
        scanf ("%c", &c);
        index++;
        if (5 == index) {
            if ('w' == c) {
                week = true;
            } else {
                week = false;
            }
        }
    } while ('\n' != c);

    if (week) {
        if (5 == x || 6 == x) {
            printf ("53\n");
        } else {
            printf ("52\n");
        }
    } else {
        if (31 == x) {
            printf ("7\n");
        } else if (30 == x) {
            printf ("11\n");
        } else {
            printf ("12\n");
        }
    }

    return 0;
}

