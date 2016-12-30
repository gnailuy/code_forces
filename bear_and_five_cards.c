#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define CARD_NUM 5

int main (int argc, char * argv[]) {
    int in, pos = 0, t[CARD_NUM] = { 0 }, total = 0;
    int t_cnt[CARD_NUM] = { 0 };
    for (int i = 0; i < CARD_NUM; i++) {
        scanf ("%d", &in);
        total += in;
        bool is_new = true;
        for (int j = 0; j < pos; j++) {
            if (t[j] == in) {
                t_cnt[j]++;
                is_new = false;
            }
        }
        if (is_new) {
            t[pos] = in;
            t_cnt[pos++] = 1;
        }
    }

    int max = 0, max_pos = -1;
    for (int i = 0; i < pos; i++) {
        int sum = -1;
        if (2 == t_cnt[i]) {
            sum = t[i] * t_cnt[i];
        } else if (t_cnt[i] > 2) {
            sum = t[i] * 3;
        }
        if (sum > max) {
            max = sum;
            max_pos = i;
        }
    }

    printf ("%d\n", total - max);

    return 0;
}

