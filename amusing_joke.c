#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LINE 100

int readline (char * buffer, int buffer_length) {
    int index = 0;
    char trash_bin;

    do {
        scanf ("%c", &buffer[index++]);
    } while (buffer[index - 1] != '\n' && index < buffer_length);

    if (buffer[index - 1] != '\n') {
        do {
            scanf ("%c", &trash_bin);
        } while (trash_bin != '\n');
    }

    buffer[index - 1] = '\0';

    return index - 1;
}

void merge (char * buffer, int start, int middle, int end) {
    char * sorted = (char *) malloc (sizeof(char) * (end - start + 1));

    int i = start, j = middle + 1, k = 0;
    while (i <= middle && j <= end) {
        if (buffer[i] <= buffer[j]) {
            sorted[k++] = buffer[i++];
        } else {
            sorted[k++] = buffer[j++];
        }
    }
    while (i <= middle) {
        sorted[k++] = buffer[i++];
    }
    while (j <= end) {
        sorted[k++] = buffer[j++];
    }
    i = start; k = 0;
    while (i <= end) {
        buffer[i++] = sorted[k++];
    }

    free (sorted);
}

void merge_sort (char * buffer, int start, int end) {
    if (end > start) {
        int middle = (end + start) / 2;
        merge_sort (buffer, start, middle);
        merge_sort (buffer, middle + 1, end);
        merge (buffer, start, middle, end);
    }
}

int main (int argc, char * argv[]) {
    char * name1 = (char *) malloc (sizeof(char) * MAX_LINE + 1);
    char * name2 = (char *) malloc (sizeof(char) * MAX_LINE + 1);
    char * letters = (char *) malloc (sizeof(char) * MAX_LINE + 1);
    int name1_length = readline (name1, MAX_LINE + 1);
    int name2_length = readline (name2, MAX_LINE + 1);
    int letters_length = readline (letters, MAX_LINE + 1);

    bool result = true;
    if (name1_length + name2_length != letters_length) {
        result = false;
    } else {
        for (int i = 0; i < name2_length; i++) {
            name1[name1_length++] = name2[i];
        }
        name1[name1_length] = '\0';

        merge_sort (name1, 0, name1_length - 1);
        merge_sort (letters, 0, letters_length - 1);

        for (int i = 0; i < letters_length; i++) {
            if (name1[i] != letters[i]) {
                result = false;
                break;
            }
        }
    }

    if (result) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (name1);
    free (name2);
    free (letters);
    return 0;
}

