#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void swap (int * buffer, int i, int j) {
    buffer[i] ^= buffer[j];
    buffer[j] ^= buffer[i];
    buffer[i] ^= buffer[j];
}

void bubble_sort (int * buffer, int * indices, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (buffer[i] < buffer[j]) { // Inverse
                swap (buffer, i, j);
                swap (indices, i, j);
            }
        }
    }
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    int * ratings = (int *) malloc (sizeof(int) * n);
    int * indices = (int *) malloc (sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d", &ratings[i]);
        indices[i] = i + 1;
    }
    bubble_sort (ratings, indices, n);

    int * rank = (int *) malloc (sizeof(int) * n);
    int curr_rating = -1, curr_index = 1;
    for (int i = 0; i < n; i++) {
        if (ratings[i] != curr_rating) {
            rank[indices[i] - 1] = curr_index;
        } else {
            rank[indices[i] - 1] = rank[indices[i - 1] - 1];
        }

        curr_index++;
        curr_rating = ratings[i];
    }

    for (int i = 0; i < n; i++) {
        printf ("%d ", rank[i]);
    }
    printf ("\n");

    free (ratings);
    free (indices);
    free (rank);
    return 0;
}

