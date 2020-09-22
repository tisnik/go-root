/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>
#include <stdint.h>

int main(void) {
    int a[3][4];
    int i, j;

    puts("Pole pred upravou:");
    for (j=0; j<3; j++) {
        for (i=0; i<4; i++) {
            printf(" %2d", a[j][i]);
        }
        putchar('\n');
    }
    puts("\n");

    for (i=0; i<12; i++) {
        a[0][i] = i;
    }

    puts("Pole po uprave:");
    for (j=0; j<3; j++) {
        for (i=0; i<4; i++) {
            printf(" %2d", a[j][i]);
        }
        putchar('\n');
    }
    return 0;
}
