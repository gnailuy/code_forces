#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define TRIE_WIDTH (26 + 10) // Lower case letters + digits
#define MAX_INPUT 32

struct trie_node_st {
    bool is_end;
    int index;
    struct trie_node_st * next[TRIE_WIDTH];
};

// a-z -> 0-25; 0-9 -> 26-35
int char_to_index (char in) {
    if ('a' <= in && in <= 'z') {
        return in - 'a';
    } else if ('0' <= in && in <= '9') {
        return in - '0' + 26;
    }
    return -1;
}

// 0-25 -> a-z; 26-35 -> 0-9
char index_to_char (int idx) {
    if (0 <= idx && idx <= 25) {
        return 'a' + idx;
    } else if (26 <= idx && idx <= 35) {
        return '0' + idx - 26;
    }
    return 'X';
}

void add_to_trie (struct trie_node_st * root, char * string, int length) {
    struct trie_node_st * head = root;
    for (int i = 0; i < length && string[i] != '\n'; i++) {
        int index = char_to_index (string[i]);
        if (head -> next[index] == NULL) {
            head -> next[index]
                = (struct trie_node_st *) malloc (sizeof(struct trie_node_st));
            head -> next[index] -> is_end = false;
            head -> next[index] -> index = 0;
            for (int j = 0; j < TRIE_WIDTH; j++) {
                (head -> next[index] -> next)[j] = NULL;
            }
        }
        head = head -> next[index];
    }
    head -> is_end = true;
}

struct trie_node_st * find_tail (struct trie_node_st * root, char * string, int length) {
    struct trie_node_st * tail = root;
    for (int i = 0; i < length && string[i] != '\n'; i++) {
        int index = char_to_index (string[i]);
        if (tail -> next[index] == NULL) {
            return NULL;
        }
        tail = tail -> next[index];
    }
    if (tail -> is_end) {
        return tail;
    }
    return NULL;
}

void destroy_trie (struct trie_node_st * head) {
    for (int i = 0; i < TRIE_WIDTH; i++) {
        if (head -> next[i] != NULL) {
            destroy_trie (head -> next[i]);
        }
    }
    free (head);
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    struct trie_node_st * root
        = (struct trie_node_st *) malloc (sizeof(struct trie_node_st));
    root -> is_end = false;
    root -> index = 0;
    for (int j = 0; j < TRIE_WIDTH; j++) {
        (root -> next)[j] = NULL;
    }

    char * input = (char *) malloc (sizeof(char) * MAX_INPUT + 2);
    struct trie_node_st * tail;
    for (int i = 0; i < n; i++) {
        fgets (input, MAX_INPUT + 2, stdin); // '\n' and '\0'
        tail = find_tail (root, input, MAX_INPUT + 1);
        if (NULL != tail) {
            for (int j = 0; j < MAX_INPUT + 1; j++) {
                if (input[j] == '\n') {
                    input[j] = '\0'; // End the string
                    break;
                }
            }
            printf ("%s%d\n", input, ++(tail -> index));
        } else {
            add_to_trie (root, input, MAX_INPUT + 1); // '\0'
            printf ("OK\n");
        }
    }

    free (input);
    destroy_trie (root);
    return 0;
}

