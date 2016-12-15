#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n;
    char in_char = ' ';
    scanf ("%d\n", &n);

    int anton = 0, danik = 0;
    for (int i = 0; i < n && in_char != '\n'; i++) {
        scanf ("%c", &in_char);
        if ('A' == in_char) anton++;
        else if ('D' == in_char) danik++;
    }

    if (anton > danik) {
        printf ("Anton\n");
    } else if (anton < danik) {
        printf ("Danik\n");
    } else {
        printf ("Friendship\n");
    }

    return 0;
}

