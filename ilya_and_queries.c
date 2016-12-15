#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_LENGTH 100000

// Find the i such that buffer[i - 1] < target && buffer[i] >= target
// Return -1 if such i does not exist
int binary_search (int * buffer, int start, int end, int target) {
    if (start > end) {
        return -1;
    }

    int middle = (start + end) / 2;
    if (buffer[middle] >= target && (middle < 1 || buffer[middle - 1] < target)) {
        return middle;
    } else if (buffer[middle] > target) {
        return binary_search (buffer, start, middle - 1, target);
    } else {
        return binary_search (buffer, middle + 1, end, target);
    }
}

int main (int argc, char * argv[]) {
    char * s = (char *) malloc (sizeof(char) * (MAX_LENGTH + 2));
    fgets (s, MAX_LENGTH + 2, stdin);

    int * is = (int *) malloc (sizeof(int) * (MAX_LENGTH - 1));
    int is_count = 0;
    for (int i = 0; i < MAX_LENGTH - 1 && s[i] != '\n'; i++) {
        if (s[i] == s[i + 1]) {
            is[is_count] = i + 1;
            is_count += 1;
        }
    }

    int m, l, r;
    scanf ("%d", &m);
    for (int i = 0; i < m; i++) {
        scanf ("%d %d", &l, &r);
        int start = binary_search (is, 0, is_count - 1, l);
        int end = binary_search (is, 0, is_count - 1, r);
        int count = 0;
        if (start >= 0) {
            if (end < 0) {
                count = is_count - start;
            } else {
                count = end - start;
            }
        }
        printf ("%d\n", count);
    }

    free (s);
    free (is);
    return 0;
}

