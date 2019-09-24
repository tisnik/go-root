# Parametry prekladace.
CFLAGS=-Wall -ansi

all:	01_boolean


# potentially dangerous thing!!!
clean:
	find . -type f -executable -delete


%:	%.c
	$(CC) $(CFLAGS) $< -o $@
