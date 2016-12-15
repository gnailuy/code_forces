#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_BUFFER 30

bool is_digit (char c) {
    return c >= '0' && c <= '9';
}

void number_to_buffer (int num, char * buffer) {
    int div = num / 26;
    int remain = num % 26;
    if (0 == remain) {
        --div;
    }

    if (div > 0) {
        number_to_buffer (div, buffer);
    }

    int i;
    for (i = 0; i < MAX_BUFFER && buffer[i] != '\0'; i++) {}
    if (remain == 0) {
        buffer[i] = 'Z';
    } else {
        buffer[i] = 'A' + (remain - 1);
    }
    buffer[i + 1] = '\0';
}

void to_excel (char * coord, char * output) {
    int y = 0;
    bool at_y = false;
    for (int i = 0; i < MAX_BUFFER && coord[i] != '\n'; i++) {
        if ('C' == coord[i]) {
            at_y = true;
            continue;
        }
        if (at_y) {
            y = y * 10 + (coord[i] - '0');
        }
    }

    output[0] = '\0';
    number_to_buffer (y, output);

    int pos = 0;
    for (pos = 0; pos < MAX_BUFFER && output[pos] != '\0'; pos++) {}

    for (int i = 1; i < MAX_BUFFER && coord[i] != 'C'; i++) {
        output[pos++] = coord[i];
    }
    output[pos] = '\0';
}

void decimal_number_to_buffer (int num, char * buffer) {
    int div = num / 10;
    int remain = num % 10;

    if (div > 0) {
        decimal_number_to_buffer (div, buffer);
    }

    int i;
    for (i = 0; i < MAX_BUFFER && buffer[i] != '\0'; i++) {}
    buffer[i] = '0' + remain;
    buffer[i + 1] = '\0';
}

void to_rxcy (char * coord, char * output) {
    int y = 0, i;
    for (i = 0; i < MAX_BUFFER && !is_digit(coord[i]); i++) {
        y = y * 26 + coord[i] - 'A' + 1;
    }

    int j = 0;
    output[j++] = 'R';
    for (; i < MAX_BUFFER && coord[i] != '\n'; i++) {
        output[j++] = coord[i];
    }
    output[j++] = 'C';
    output[j] = '\0';

    decimal_number_to_buffer (y, output);
    int pos = 0;
    for (pos = 0; pos < MAX_BUFFER && output[pos] != '\0'; pos++) {}
    output[pos] = '\0';
}

int main (int argc, char * argv[]) {
    int n;
    scanf ("%d\n", &n);

    char * coord = (char *) malloc (sizeof(char) * MAX_BUFFER);
    char * another = (char *) malloc (sizeof(char) * MAX_BUFFER);

    for (int i = 0; i < n; i++) {
        fgets (coord, MAX_BUFFER, stdin);
        bool is_rxcy = false;
        if ('R' == coord[0] && is_digit(coord[1])) {
            for (int j = 0; j < MAX_BUFFER && coord[j] != '\n'; j++) {
                if ('C' == coord[j]) {
                    is_rxcy = true;
                    break;
                }
            }
        }

        if (is_rxcy) {
            to_excel (coord, another);
        } else {
            to_rxcy (coord, another);
        }
        printf ("%s\n", another);
    }

    free (coord);
    free (another);
    return 0;
}

