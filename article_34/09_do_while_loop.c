#include <stdio.h>

int main(void)
{
    int x = 1;

    do {
        printf("%d\n", x);
        x++;
    } while (x <= 10);

    return 0;
}

