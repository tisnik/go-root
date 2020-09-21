/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

enum {
    Pondeli,
    Utery,
    Streda,
    Ctvrtek,
    Patek,
    Sobota,
    Nedele
};

int main(void) {
    printf("%d\n", Pondeli);
    printf("%d\n", Streda);
    printf("%d\n", Patek);
    return 0;
}
