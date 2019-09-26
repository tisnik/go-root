#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    const char *s1 = "Hello ";
    const char *s2 = "world!";

    size_t length = strlen(s1) + strlen(s2) + 1;
    char *s3 = (char *)calloc(length, sizeof(char));

    if (s3) {
        strcpy(s3, s1);
        strcat(s3, s2);

        puts(s3);
        free(s3);
    }
    else {
        perror("calloc() failed");
    }

    return 0;
}
