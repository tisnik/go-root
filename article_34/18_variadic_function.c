#include <stdio.h>
#include <stdarg.h>

void f1(const char *msg) {
    printf("%s\n", msg);
}

void f2(int count, ...) {
    int i;
    va_list args;
    va_start(args, count);

    for (i = 0; i < count; i++) {
        char *msg = va_arg(args, char *);
        printf("%s ", msg);
    }
    putchar('\n');
    va_end(args);
}

void f3(char *prefix, int count, ...) {
    int i;
    va_list args;

    printf("%s ", prefix);
    va_start(args, count);

    for (i = 0; i < count; i++) {
        char *msg = va_arg(args, char *);
        printf("%s ", msg);
    }
    putchar('\n');
    va_end(args);
}

int main(void) {
    f1("Hello");
    f2(3, "Hello", "world", "!");
    f3("Message:", 4, "Hello", "world", "again", "!");
    return 0;
}
