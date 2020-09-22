/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>

int main(void) {
    const char *s = "www.root.cz";

    puts(s);
    s[3] = '*';
    s[8] = '*';
    puts(s);

    return 0;
}
