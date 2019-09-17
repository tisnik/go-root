# Parametry prekladace.
CFLAGS=-Wall -ansi

all:	01_empty_main 02_hello_world 03_data_types 04_if_then 05_if_then_else \
	06_nested_if 07_while_loop 08_while_loop 09_do_while_loop \
	10_for_loop 11_for_loop  12_for_loop 13_break_in_loop 14_switch


# potentially dangerous thing!!!
clean:
	find . -type f -executable -delete


%:	%.c
	$(CC) $(CFLAGS) $< -o $@
