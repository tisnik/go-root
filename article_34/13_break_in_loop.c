#include <stdio.h>

int main(void)
{
    int x = 1;

    while (1) { /* endless loop */
        printf("%d\n", x);
        x++;
        if (x > 10) break;
    }

    return 0;
}

