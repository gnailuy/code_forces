#include <stdio.h>
#include <stdlib.h>

int distance (int n, int start, int end) {
    if (start <= end) {
        return end - start;
    } else {
        return end + n - start;
    }
}

int main (int argc, char * argv[]) {
    int n, m;
    long long time = 0;
    scanf ("%d %d", &n, &m);

    int * tasks = malloc (sizeof(int) * (m + 1));
    tasks[0] = 1; // Start from house #1
    for (int i = 1; i <= m; i++) {
        scanf ("%d", &tasks[i]);
    }

    for (int i = 0; i < m; i++) {
        time += distance (n, tasks[i], tasks[i + 1]);
    }

    // printf ("%I64d\n", time);
    printf ("%lld\n", time);

    free (tasks);
    return 0;
}

