#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, mishka, chris, mishka_count = 0, chris_count = 0;
    scanf ("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf ("%d %d", &mishka, &chris);
        if (mishka > chris) {
            mishka_count++;
        } else if (mishka < chris) {
            chris_count++;
        }
    }

    if (mishka_count > chris_count) {
        printf ("Mishka\n");
    } else if (mishka_count < chris_count) {
        printf ("Chris\n");
    } else {
        printf ("Friendship is magic!^^\n");
    }

    return 0;
}

