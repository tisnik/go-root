# Parametry prekladace.
CFLAGS=-Wall -ansi

all:	01_booleans 03_strings 04_immutable_strings 05_mutable_strings 06_mutable_utf8_strings \
	07_string_comparison 08_string_concatenation


# potentially dangerous thing!!!
clean:
	find . -type f -executable -delete


%:	%.c
	$(CC) $(CFLAGS) $< -o $@
