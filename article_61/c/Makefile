# Parametry prekladace.
CC=gcc
CFLAGS=-Wall -ansi -pedantic
LDFLAGS=-lSDL2

all:	test1 test2 test3 test4 fractals

clean:
	rm -f test1 test2 test3 test4 fractals

# Pravidlo pro preklad a slinkovani prislusneho souboru
# do vysledne spustitelne aplikace.
%:	%.c
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS)

fractals:	fractals.c
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS) -lm
