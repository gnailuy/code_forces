#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct DDates {
    int exam_date;
    int early_date;
} DDates;

int compare (DDates a, DDates b) {
    if (a.exam_date > b.exam_date) {
        return 1;
    } else if (a.exam_date < b.exam_date) {
        return -1;
    } else if (a.early_date == b.early_date) {
        return 0;
    } else {
        return a.early_date > b.early_date ? 1 : -1;
    }
}

void swap (DDates * buffer, int i, int j) {
    buffer[i].exam_date ^= buffer[j].exam_date;
    buffer[j].exam_date ^= buffer[i].exam_date;
    buffer[i].exam_date ^= buffer[j].exam_date;

    buffer[i].early_date ^= buffer[j].early_date;
    buffer[j].early_date ^= buffer[i].early_date;
    buffer[i].early_date ^= buffer[j].early_date;
}

void bubble_sort (DDates * buffer, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (compare (buffer[i], buffer[j]) > 0) {
                swap (buffer, i, j);
            }
        }
    }
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    DDates * exams = (DDates *) malloc (sizeof(DDates) * n);
    for (int i = 0; i < n; i++) {
        scanf ("%d %d", &(exams[i].exam_date), &(exams[i].early_date));
    }

    bubble_sort (exams, n);

    int current_date = 0;
    for (int i = 0; i < n; i++) {
        if (exams[i].early_date >= current_date) {
            current_date = exams[i].early_date;
        } else {
            current_date = exams[i].exam_date;
        }
    }

    printf ("%d\n", current_date);

    free (exams);
    return 0;
}

