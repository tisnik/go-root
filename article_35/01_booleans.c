/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>

int main(void) {
    if (NULL) {
        puts("true");
    }

    if (!NULL) {
        puts("false");
    }

    if (1) {
        puts("true");
    }

    if (0) {
        puts("false");
    }

    if (1.5) {
        puts("true");
    }

    if (0.0) {
        puts("false");
    }

    void *b1 = NULL;

    if (b1) {
        puts("true");
    }

    if (!b1) {
        puts("false");
    }

    return 0;
}
