#include <stdio.h>

int main(void) {
    int x = 1;
    int *px = &x;
    (*px)++;
    printf("%d\n", x);
    return 0;
}

