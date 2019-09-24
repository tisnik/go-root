#include <stdio.h>

int main(void) {
    char s[] = "www.root.cz";

    puts(s);
    s[3] = '*';
    s[8] = '*';
    puts(s);

    return 0;
}
