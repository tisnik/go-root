/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

int main(void)
{
    int i,x;

    for (i=0, x=1; i <= 10; i++, x<<=1) {
        printf("2^%d = %d\n", i, x);
    }

    return 0;
}

