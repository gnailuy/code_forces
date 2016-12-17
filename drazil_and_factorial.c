#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void swap (int * buffer, int i, int j) {
    buffer[i] ^= buffer[j];
    buffer[j] ^= buffer[i];
    buffer[i] ^= buffer[j];
}

void bubble_sort (int * buffer, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (buffer[i] < buffer[j]) { // Reverse
                swap (buffer, i, j);
            }
        }
    }
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    int * digits = (int *) malloc (sizeof(int) * n * 4);
    int dp = 0; // Pointer

    char in;
    for (int i = 0; i < n; i++) {
        scanf ("%c", &in);
        if (in == '9') {
            digits[dp++] = 7;
            digits[dp++] = 3;
            digits[dp++] = 3;
            digits[dp++] = 2;
        } else if (in == '8') {
            digits[dp++] = 7;
            digits[dp++] = 2;
            digits[dp++] = 2;
            digits[dp++] = 2;
        } else if (in == '6') {
            digits[dp++] = 5;
            digits[dp++] = 3;
        } else if (in == '4') {
            digits[dp++] = 3;
            digits[dp++] = 2;
            digits[dp++] = 2;
        } else if (in != '1' && in != '0') {
            digits[dp++] = in - '0';
        }
    }
    bubble_sort (digits, dp);

    for (int i = 0; i < dp; i++) {
        printf ("%d", digits[i]);
    }
    printf ("\n");

    return 0;
}

