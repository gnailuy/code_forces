#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char * argv[]) {
    int n, m;
    scanf ("%d %d\n", &n, &m);
    int line_chars = m * 2 + 1;

    char * line = (char *) malloc (sizeof(char) * line_chars);
    bool is_colored = false;
    for (int j = 0; j < n; j++) {
        fgets(line, line_chars, stdin);
        if (!is_colored) {
            for (int i = 0; i < line_chars; i++) {
                if ('C' == line[i] || 'M' == line[i] || 'Y' == line[i]) {
                    is_colored = true;
                }
            }
        }
    }

    if (is_colored) {
        printf ("#Color\n");
    } else {
        printf ("#Black&White\n");
    }

    free (line);
    return 0;
}

