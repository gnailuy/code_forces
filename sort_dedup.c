#include <stdio.h>
#include <stdlib.h>

void swap(int * array, int i, int j) {
    if (i == j) return;

    int tmp = array[i];
    array[i] = array[j];
    array[j] = tmp;
}

void bubble_sort(int * array, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i+1; j < length; j++) {
            if (array[i] > array[j]) {
                swap(array, i, j);
            }
        }
    }
}

int main(int argc, char * argv[]) {
    // Read file, TODO: read filename from argv
    FILE * ifp = fopen("input.txt", "r");
    if (!ifp) {
        printf("Input file pointer error\n");
    }

    // Read length, the first number in the file
    int length = -1;
    if (fscanf(ifp, "%d", &length) == EOF || length <= 0) {
        printf("Length error: %d\n", length);
    }

    // Read array, the second line of the file
    int * array = malloc(sizeof(int) * length);
    for (int i = 0; i < length; i++) {
        if (fscanf(ifp, "%d", &array[i]) == EOF) {
            printf("Read error\n");
        }
    }

    // Sort the array, TODO: try merge_sort or quick_sort
    bubble_sort(array, length);

    // Dedup
    int tail_index = 0;
    for (int i = 1; i < length; i++) {
        if (array[i] != array[tail_index]) {
            swap(array, i, ++tail_index);
        }
    }

    // Open file for output
    FILE * ofp = fopen("output.txt", "w+");
    if (!ofp) {
        printf("Output file pointer error\n");
    }

    fprintf(ofp, "%d\n", tail_index+1);
    for (int i = 0; i <= tail_index; i++) {
        fprintf(ofp, "%d ", array[i]);
    }
    fprintf(ofp, "\n");
}

