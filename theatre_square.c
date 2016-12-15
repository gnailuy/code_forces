#include <stdio.h>
#include <stdlib.h>

unsigned long long get_square_number (int m, int a) {
    int m_num = m / a;
    if (m_num < (m / (a + 0.0))) {
        return m_num + 1;
    } else {
        return m_num;
    }
}

int main (int argc, char * argv[]) {
    int m, n, a;

    scanf ("%d", &m);
    scanf ("%d", &n);
    scanf ("%d", &a);

    printf ("%llu\n", get_square_number(m, a) * get_square_number(n, a));
    return 0;
}
