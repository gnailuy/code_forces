#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_NUM 100000

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    int * a = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &a[i]);
    }

    int * nums = (int *) calloc (MAX_NUM, sizeof(int)); // Zero-initialized
    int * suffixes = (int *) malloc (sizeof(int) * n);
    nums[a[n - 1] - 1] = 1;
    suffixes[n - 1] = 1;
    for (int i = n - 2; i >= 0; i--) {
        if (0 == nums[a[i] - 1]) {
            nums[a[i] - 1] = 1;
            suffixes[i] = suffixes[i + 1] + 1;
        } else {
            suffixes[i] = suffixes[i + 1];
        }
    }

    int l;
    for (int i = 0; i < m; i++) {
        scanf ("%d", &l);
        printf ("%d\n", suffixes[l - 1]);
    }

    free (a);
    free (nums);
    return 0;
}

