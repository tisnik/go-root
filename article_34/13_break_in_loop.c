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

    while (1) { /* endless loop */
        printf("%d\n", x);
        x++;
        if (x > 10) break;
    }

    return 0;
}

