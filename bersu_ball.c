#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_SKILL 100

int pair_bng (int *a, int *b) {
    int pairs = 0;

    if (*a >= *b) {
        pairs += *b;
        *a -= *b;
        *b = 0;
    } else {
        pairs += *a;
        *b -= *a;
        *a = 0;
    }

    return pairs;
}

int main (int argc, char * argv[]) {
    int n, m, skill;

    scanf ("%d", &n);
    int * boys = (int *) calloc (MAX_SKILL, sizeof(int)); // Zero-initialized
    for (int i = 0; i < n; i++) {
        scanf ("%d", &skill);
        boys[skill - 1]++;
    }

    scanf ("%d", &m);
    int * girls = (int *) calloc (MAX_SKILL, sizeof(int)); // Zero-initialized
    for (int i = 0; i < m; i++) {
        scanf ("%d", &skill);
        girls[skill - 1]++;
    }

    int pairs = 0;
    for (int i = 0; i < MAX_SKILL; i++) {
        // Pair with b&gs at the last skill
        if (i > 0 && boys[i - 1] > 0 && girls[i] > 0) {
            pairs += pair_bng (&boys[i - 1], &girls[i]);
        } else if (i > 0 && girls[i - 1] > 0 && boys[i] > 0) {
            pairs += pair_bng (&girls[i - 1], &boys[i]);
        }

        // Pair this skill
        if (boys[i] > 0 && girls[i] > 0) {
            pairs += pair_bng (&girls[i], &boys[i]);
        }
    }

    printf ("%d\n", pairs);

    return 0;
}

