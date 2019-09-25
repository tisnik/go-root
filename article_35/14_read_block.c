#include <stdio.h>
#include <stdint.h>

const char* filename = "test_input.txt";
const int BLOCK_SIZE = 16;

int main(void) {
    FILE *fin = fopen(filename, "r");
    if (!fin) {
        perror("Can not open file");
        return 1;
    }

    char buffer[BLOCK_SIZE];
    while (1) {
        size_t read = fread(buffer, sizeof(char), BLOCK_SIZE, fin);
        if (read > 0) {
            int i;
            printf("read %ld bytes\n", read);
            for (i=0; i<read; i++) {
                printf("%d ", buffer[i]);
            }
            putchar('\n');
        }
        else {
            puts("reached end of file");
            break;
        }
    }
    fclose(fin);

    return 0;
}
