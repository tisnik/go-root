/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

int main(void)
{
    int x = 10;

    if (x > 0) {
        puts("x is positive number");
    }
    else {
        puts("x is negative number or zero");
    }

    return 0;
}
