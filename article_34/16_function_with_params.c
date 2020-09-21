/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

void printHello(const char *message) {
    puts(message);
}

int main(void) {
    printHello("Hello world!");
    return 0;
}
