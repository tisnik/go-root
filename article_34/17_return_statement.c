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
