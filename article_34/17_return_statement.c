/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

const char *getMessage(void) {
    return "Hello world!";
}

void printMessage(const char *message) {
    puts(message);
}

int main(void) {
    printMessage(getMessage());
    return 0;
}
