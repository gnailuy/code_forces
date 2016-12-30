#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

char * digits[10] = { // 0 ~ 9
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine"
};

char * teens[10] = { // 10 ~ 19
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fifteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen"
};

char * tys[10] = { // 0, 10, 20 ~ 90
    "zero",
    "ten",
    "twenty",
    "thirty",
    "forty",
    "fifty",
    "sixty",
    "seventy",
    "eighty",
    "ninety"
};

int main (int argc, char * argv[]) {
    int s;
    scanf ("%d", &s);

    if (s < 10) {
        printf ("%s\n", digits[s]);
    } else if (s < 20) {
        printf ("%s\n", teens[s - 10]);
    } else if (s % 10 == 0) {
        printf ("%s\n", tys[s / 10]);
    } else {
        printf ("%s-%s\n", tys[s / 10], digits[s % 10]);
    }

    return 0;
}

