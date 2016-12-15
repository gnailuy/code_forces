#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

#define MAX_LEN 10
#define BUFFER_LEN (MAX_LEN * 2 + 3)
#define WORD_BUFFER (MAX_LEN + 1)

#define LETTERS 26
#define MAX_STORAGE 3000

typedef struct node {
    struct node * letters[LETTERS];
    int index;
} f_node;

f_node * new_node (int index) {
    f_node * res = (f_node *) malloc (sizeof(f_node));
    for (int i = 0; i < LETTERS; i++) {
        res -> letters[i] = NULL;
    }
    res -> index = index;

    return res;
}

char * new_word () {
    return (char *) malloc (sizeof(char) * WORD_BUFFER);
}

void insert (f_node * forest, char ** storage, int index,
             char * key, char * value) { // Both key and value are tailing '\0'
    f_node * pointer = forest;
    int pos = 0;
    char k;
    while (true) {
        k = key[pos++];
        if ('\0' != k) {
            if (NULL == pointer -> letters[k - 'a']) {
                pointer -> letters[k - 'a'] = new_node (-1);
            }
            pointer = pointer -> letters[k - 'a'];
        } else {
            pointer -> index = index;
            break;
        }
    }

    storage[index] = new_word ();
    for (int i = 0; i < WORD_BUFFER + 1; i++) {
        storage[index][i] = value[i];
        if ('\0' == value[i]) {
            break;
        }
    }
}

int find_index (f_node * forest, char * key) {
    f_node * pointer = forest;
    int pos = 0;
    while ('\0' != key[pos]) {
        pointer = pointer -> letters[key[pos++] - 'a'];
    }

    return pointer -> index;
}

void destory (f_node * forest) {
    f_node * pointer = forest;
    for (int i = 0; i < LETTERS; i++) {
        if (NULL != pointer -> letters[i]) {
            destory (pointer -> letters[i]);
        }
    }
    free (pointer);
}

void release (char ** storage, int count) {
    for (int i = 0; i < count; i++) {
        free (storage[i]);
    }
    free (storage);
}

int main(int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d\n", &n, &m);

    f_node * forest = new_node (-1);
    char ** storage = (char **) malloc (sizeof(char *) * MAX_STORAGE);
    int storage_index = 0;

    char * buffer = (char *) malloc (sizeof(char) * BUFFER_LEN); // ' ', '\n', '\0'
    char * second;
    for (int i = 0; i < m; i++) {
        fgets(buffer, BUFFER_LEN, stdin);
        for (int j = 0; j < BUFFER_LEN; j++) {
            if (' ' == buffer[j]) {
                buffer[j] = '\0';
                second = &buffer[j + 1];
            } else if ('\n' == buffer[j]) {
                buffer[j] = '\0';
                break;
            }
        }
        insert (forest, storage, storage_index++, buffer, second);
    }

    char * lec_buffer = (char *) malloc (sizeof(char) * WORD_BUFFER); // '\0'
    char in; int idx = 0, lec_length = 0;
    while (true) {
        scanf ("%c", &in);
        if (' ' == in || '\n' == in) {
            lec_buffer[idx] = '\0';
            char * s = storage[find_index(forest, lec_buffer)];
            if (lec_length > strlen (s)) {
                printf ("%s", s);
            } else {
                printf ("%s", lec_buffer);
            }

            if ('\n' == in) {
                printf ("\n");
                break;
            } else {
                printf (" ");
            }

            idx = 0;
            lec_length = 0;
        } else {
            lec_buffer[idx++] = in;
            lec_length++;
        }
    }

    destory (forest);
    release (storage, storage_index);
    free (buffer);
    return 0;
}

