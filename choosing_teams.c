#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_PARTICIPATE 5
#define TEAM_MEMBERS 3

int main (int argc, char * argv[]) {
    int n, k, participates, ok_members = 0;
    scanf ("%d %d", &n, &k);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &participates);
        if (participates <= MAX_PARTICIPATE - k) {
            ok_members++;
        }
    }

    printf ("%d\n", ok_members / TEAM_MEMBERS);

    return 0;
}

