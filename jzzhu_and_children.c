#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct Node {
    int want;
    int index;
    struct Node * next;
} LNode;

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d", &n, &m);

    LNode * head = NULL;
    LNode * tail = NULL;
    for (int i = 0; i < n; i++) {
        LNode * new_node = (LNode *) malloc (sizeof(LNode));
        scanf ("%d", &(new_node -> want));
        new_node -> index = i + 1;
        new_node -> next = NULL;

        if (NULL == head) {
            head = tail = new_node;
        } else {
            tail -> next = new_node;
            tail = tail -> next;
        }
    }

    LNode * ptr;
    while (NULL != head -> next) {
        if (m >= head -> want) {
            ptr = head;
            head = head -> next;
            free (ptr);
        } else {
            head -> want -= m;
            tail -> next = head;
            tail = tail -> next;
            head = head -> next;
            tail -> next = NULL;
        }
    }

    printf ("%d\n", head -> index);
    free (head);

    return 0;
}

