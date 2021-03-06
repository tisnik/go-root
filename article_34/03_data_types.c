/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

int main(void)
{
    printf("sizeof char        = %lu byte(s)\n", sizeof(char));
    printf("sizeof short       = %lu byte(s)\n", sizeof(short));
    printf("sizeof int         = %lu byte(s)\n", sizeof(int));
    printf("sizeof long        = %lu byte(s)\n", sizeof(long));
    printf("sizeof long long   = %lu byte(s)\n", sizeof(long long));
    printf("sizeof float       = %lu byte(s)\n", sizeof(float));
    printf("sizeof double      = %lu byte(s)\n", sizeof(double));
    printf("sizeof long double = %lu byte(s)\n", sizeof(long double));
    return 0;
}

