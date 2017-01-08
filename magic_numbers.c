#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_LENGTH 11

int int_2_str (char * buffer, int n, int buf_len) {
    buffer[buf_len - 1] = '\0';
    int i = buf_len - 2;
    while (n > 0 && i >= 0) {
        buffer[i--] = '0' + n % 10;
        n /= 10;
    }
    if (n > 0) {
        return -1;
    } else {
        return buf_len - i - 2;
    }
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d", &n);

    char * buffer = (char *) malloc (sizeof(char) * MAX_LENGTH);
    int n_len = int_2_str (buffer, n, MAX_LENGTH);
    if (n_len < 0) return -1;
    char * n_str = &buffer[MAX_LENGTH - n_len - 1];

    bool is_magic = true;
    char status = 'S';
    for (int i = 0; i < n_len; i++) {
        switch (n_str[i]) {
            case '1':
                status = 'A';
                break;
            case '4':
                if ('A' == status) {
                    status = 'B';
                } else if ('B' == status) {
                    status = 'C';
                } else {
                    is_magic = false;
                }
                break;
            default:
                is_magic = false;
                break;
        }
        if (!is_magic) break;
    }

    if (is_magic) {
        printf ("YES\n");
    } else {
        printf ("NO\n");
    }

    free (buffer);
    return 0;
}

