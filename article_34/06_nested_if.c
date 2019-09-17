#include <stdio.h>

int main(void)
{
    int x = 0;

    if (x > 0) {
        puts("x is positive number");
    }
    else if (x == 0) {
        puts("x is zero");
    }
    else {
        puts("x is negative number");
    }

    return 0;
}

