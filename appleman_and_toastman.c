#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void swap (int * buffer, int i, int j) {
    buffer[i] ^= buffer[j];
    buffer[j] ^= buffer[i];
    buffer[i] ^= buffer[j];
}

int find_parent_index (int index) {
    return (index - 1) / 2;
}

int find_left_child_index (int index) {
    return 2 * index + 1;
}

void insert_to_binary_heap (int * heap, int tail, int number) {
    heap[tail] = number;

    int parent;
    do {
        parent = find_parent_index (tail);
        if (heap[parent] > heap[tail]) {
            swap (heap, parent, tail);
            tail = parent;
        } else {
            break;
        }
    } while (parent >= 0);
}

int remove_from_binary_heap (int * heap, int tail, int remove_pos) {
    int removed_value = heap[remove_pos];

    if (remove_pos != tail) {
        heap[remove_pos] = heap[tail--];

        while (remove_pos < tail) {
            int l = find_left_child_index (remove_pos);
            if (l < tail && heap[l] > heap[l + 1]) { // Two children
                l++;
            } else if (l > tail) {
                break;
            }
            if (heap[remove_pos] <= heap[l]) {
                break;
            } else {
                swap (heap, remove_pos, l);
                remove_pos = l;
            }
        }
    }

    return removed_value;
}

int main (int argc, char * argv[]) {
    int n, number;
    scanf ("%d", &n);

    int * binary_heap = (int *) malloc (sizeof(int) * n);

    for (int i = 0; i < n; i++) {
        scanf ("%d", &number);
        insert_to_binary_heap (binary_heap, i, number);
    }

    long long sum = 0L, min;
    int index = 1;
    for (int i = n - 1; i >= 0; i--) {
        min = (long long) remove_from_binary_heap (binary_heap, i, 0);
        if (index < n) {
            sum += min * (1 + index++);
        } else {
            sum += min * (index++);
        }
    }

    // printf ("%I64d\n", sum);
    printf ("%lld\n", sum);

    free (binary_heap);
    return 0;
}

