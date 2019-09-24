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
