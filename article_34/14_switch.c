/*
 Seriál "Programovací jazyk Go"

 Třicátá čtvrtá část
    Programovací jazyk Go pro skalní céčkaře
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
*/

#include <stdio.h>

const char *classify(int x) {
	switch (x) {
	case 0:
		return "nula";
        case 2:
        case 4:
        case 6:
        case 8:
		return "sudé číslo";
	case 1:
	case 3:
	case 5:
	case 7:
	case 9:
		return "liché číslo";
	default:
		return "?";
	}
}

int main(void) {
    int x;
    for (x = 0; x <= 10; x++) {
        printf("%d: %s\n", x, classify(x));
    }
    return 0;
}
