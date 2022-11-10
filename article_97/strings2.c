#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
	char *s = (char*)malloc(100*sizeof(char));

	strcpy(s, "Hello*");
	strcat(s, "world!");
	puts(s);

	s[5] = ' ';
	puts(s);

	free(s);

}
