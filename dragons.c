#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int s, n;
    scanf ("%d %d", &s, &n);

    int best_move, best_bouns;

    int * blood = malloc (sizeof(int) * n);
    int * bonus = malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d %d", &blood[i], &bonus[i]);
    }

    bool can_pass = true;
    for (int i = 0; i < n; i++) {
        best_move = -1;
        best_bouns = -1;
        for (int j = 0; j < n; j++) {
            if (s > blood[j] && blood[j] != 0) {
                if (bonus[j] > best_bouns) {
                    best_move = j;
                    best_bouns = bonus[j];
                }
            }
        }
        if (best_move >= 0) {
            blood[best_move] = 0;
            s += best_bouns;
        } else {
            can_pass = false;
            break;
        }
    }

    if (can_pass) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (blood);
    free (bonus);
    return 0;
}

