#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int abs(int number) {
    return number > 0 ? number : -number;
}

int main (int argc, char * argv[]) {
    int count, distance, result = 0;
    scanf ("%d\n", &count);

    char * original_stat = (char *) malloc (sizeof(char) * (count + 2));
    fgets(original_stat, count + 2, stdin);

    char secret_digit;
    for (int i = 0; i < count; i++) {
        scanf ("%c", &secret_digit);
        distance = abs (secret_digit - original_stat[i]);
        result += distance <= 5 ? distance : 10 - distance;
    }
    printf ("%d\n", result);

    return 0;
}
