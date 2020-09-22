/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>

int main(void) {
    char s[] = "[шщэюя]";

    puts(s);

    s[0] = '*';

    /* problematicka cast */
    s[5] = '-';

    s[11] = '*';
    puts(s);

    return 0;
}
