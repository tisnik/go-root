/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

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

