/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>
#include <stdint.h>

int main(void) {
	int32_t a[10];
        int i;

	puts("Pole pred upravou:");
        for (i=0; i<10; i++) {
            printf(" %d", a[i]);
        }
        puts("\n");

        for (i=0; i<10; i++) {
            a[i] = i * 2;
        }

	puts("Pole po uprave:");
        for (i=0; i<10; i++) {
            printf(" %d", a[i]);
        }
        return 0;
}
