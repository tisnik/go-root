#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    const char *s1 = "Hello world!";
    char *s2 = (char*)calloc(4 + 2, sizeof(char*));

    strncpy(s2, s1 + 0, 4);

    puts(s2);

    return 0;
}
