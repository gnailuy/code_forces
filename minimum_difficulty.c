#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_DIFF 1000

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int first, second, next, max_diff_orig, min_diff_after = MAX_DIFF;
    scanf ("%d", &first);
    scanf ("%d", &second);
    max_diff_orig = second - first;
    for (int i = 2; i < n; i++) {
        scanf ("%d", &next);

        if (max_diff_orig < next - second) {
            max_diff_orig = next - second;
        }
        if (min_diff_after > next - first) {
            min_diff_after = next - first;
        }

        first = second;
        second = next;
    }

    printf ("%d\n", min_diff_after > max_diff_orig ?
                    min_diff_after : max_diff_orig);

    return 0;
}
