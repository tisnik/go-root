/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

int main(void)
{
    int x;

    for (x=1; x <= 10000; x<<=1) {
        printf("%d\n", x);
    }

    return 0;
}

