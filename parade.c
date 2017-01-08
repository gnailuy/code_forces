#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int abs (int n) {
    return n > 0 ? n : -n;
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int l_cnt = 0, r_cnt = 0;
    int max_lr_diff = 0, max_rl_diff = 0, lr_idx = -1, rl_idx = -1;
    int left, right;
    for (int i = 0; i < n; i++) {
        scanf ("%d %d", &left, &right);
        l_cnt += left;
        r_cnt += right;
        if (left > right && left - right > max_lr_diff) {
            max_lr_diff = left - right;
            lr_idx = i + 1;
        } else if (left < right && right - left > max_rl_diff) {
            max_rl_diff = right - left;
            rl_idx = i + 1;
        }
    }

    int o_score = abs (l_cnt - r_cnt);
    int a_score = abs ((l_cnt - max_lr_diff) - (r_cnt + max_lr_diff));
    int b_score = abs ((l_cnt + max_rl_diff) - (r_cnt - max_rl_diff));

    if (a_score > o_score && b_score <= o_score) {
        printf ("%d\n", lr_idx);
    } else if (a_score <= o_score && b_score > o_score) {
        printf ("%d\n", rl_idx);
    } else if (a_score > o_score && b_score > o_score) {
        if (a_score >= b_score) {
            printf ("%d\n", lr_idx);
        } else {
            printf ("%d\n", rl_idx);
        }
    } else if (a_score <= o_score && b_score <= o_score) {
        printf ("0\n");
    }

    return 0;
}

