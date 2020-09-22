/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>
#include <stdint.h>

const char* filename = "test_input.txt";

int main(void) {
    FILE *fin = fopen(filename, "r");
    if (!fin) {
        perror("Can not open file");
        return 1;
    }
    char c;
    while ((c=fgetc(fin)) != EOF) {
        printf("%c", c);
    }
    fclose(fin);

    return 0;
}
