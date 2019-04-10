package main

import (
	"github.com/c-bata/go-prompt"
	"os"
)

var suggestions = []prompt.Suggest{
	{Text: "ABS", Description: "Returns the absolute value of a number"},
	{Text: "ADR", Description: "Returns the address in memory of a variable (mostly used for machine code routines stored in variables)"},
	{Text: "AND", Description: "Logical conjunction"},
	{Text: "ASC", Description: "Returns the ATASCII value of a character"},
	{Text: "ATN", Description: "Returns the arctangent of a number"},
	{Text: "BYE", Description: "Transfers control to the internal Self Test program"},
	{Text: "CHR", Description: "Returns a character given an ATASCII value"},
	{Text: "CLOAD", Description: "Loads from cassette tape a tokenized program that was saved with CSAVE"},
	{Text: "CLOG", Description: "Returns the common logarithm of a number"},
	{Text: "CLOSE", Description: "Terminates pending transfers (flush) and closes an I/O channel"},
	{Text: "CLR", Description: "Clears variables' memory and program stack"},
	{Text: "COLOR", Description: "Chooses which logical color to draw in"},
	{Text: "COM", Description: "Implementation of MS Basic's COMMON was cancelled. Recognized but the code for DIM is executed instead"},
	{Text: "CONT", Description: "Resumes execution of a program after a STOP at the next line number (see STOP)"},
	{Text: "COS", Description: "Returns the cosine of a number"},
	{Text: "CSAVE", Description: "Saves to cassette tape a program in tokenized form (see CLOAD)"},
	{Text: "DATA", Description: "Stores data in lists of numeric or string values"},
	{Text: "DEG", Description: "Switches trigonometric functions to compute in degrees (radians is the default mode) (see RAD)"},
	{Text: "DIM", Description: "Defines the size of a string or array (see COM)"},
	{Text: "DOS", Description: "Transfers control to the Disk Operating System (DOS); if DOS was not loaded, same as BYE"},
	{Text: "DRAWTO", Description: "Draws a line to given coordinates"},
	{Text: "END", Description: "Finishes execution of the program, closes open I/O channels and stops any sound"},
	{Text: "ENTER", Description: "Loads and merges into memory a plain text program from an external device, usually from cassette tape or disk (see LIST)"},
	{Text: "EXP", Description: "Exponential function"},
	{Text: "FOR", Description: "Starts a for loop"},
	{Text: "FRE", Description: "Returns the amount of free memory in bytes"},
	{Text: "GET", Description: "Reads one byte from an I/O channel (see PUT)"},
	{Text: "GOSUB", Description: "Jumps to a subroutine at a given line in the program, placing the return address on the stack (see POP and RETURN)"},
	{Text: "GOTO", Description: "and GO TO 	Jumps to a given line in the program. GOTO can be omitted in IF ... THEN GOTO ..."},
	{Text: "GRAPHICS", Description: "Sets the graphics mode"},
	{Text: "IF", Description: "Executes code depending on whether a condition is true or not"},
	{Text: "INPUT", Description: "Retrieves a stream of text from an I/O channel; usually data from keyboard (default), cassette tape or disk"},
	{Text: "INT", Description: "Returns the floor of a number"},
	{Text: "LEN", Description: "Returns the length of a string"},
	{Text: "LET", Description: "Assigns a value to a variable. LET can be omitted"},
	{Text: "LIST", Description: "Lists (all or part of) the program to screen (default), printer, disk, cassette tape, or any other external device (see ENTER)"},
	{Text: "LOAD", Description: "Loads a tokenized program from an external device; usually a cassette tape or disk (see SAVE)"},
	{Text: "LOCATE", Description: "Stores the logical color or ATASCII character at given coordinates"},
	{Text: "LOG", Description: "Returns the natural logarithm of a number"},
	{Text: "LPRINT", Description: "Prints text to a printer device (same result can be achieved with OPEN, PRINT and CLOSE statements)"},
	{Text: "NEW", Description: "Erases the program and all the variables from memory; automatically executed before a LOAD or CLOAD"},
	{Text: "NEXT", Description: "Continues the next iteration of a FOR loop"},
	{Text: "NOT", Description: "Logical negation"},
	{Text: "NOTE", Description: "Returns the current position on an I/O channel"},
	{Text: "ON", Description: "A computed goto - performs a jump based on the value of an expression"},
	{Text: "OPEN", Description: "Initialises an I/O channel"},
	{Text: "OR", Description: "Logical disjunction"},
	{Text: "PADDLE", Description: "Returns the position of a paddle controller"},
	{Text: "PEEK", Description: "Returns the value at an address in memory"},
	{Text: "PLOT", Description: "Draws a point at given coordinates"},
	{Text: "POINT", Description: "Sets the current position on an I/O channel"},
	{Text: "POKE", Description: "Sets a value at an address in memory"},
	{Text: "POP", Description: "Removes a subroutine return address from the stack (see GOSUB and RETURN)"},
	{Text: "POSITION", Description: "Sets the position of the graphics cursor"},
	{Text: "PRINT", Description: "and ? 	Writes text to an I/O channel; usually to screen (default), printer, cassette tape or disk (see LPRINT and INPUT)"},
	{Text: "PTRIG", Description: "Indicates whether a paddle trigger is pressed or not"},
	{Text: "PUT", Description: "Writes one byte to an I/O channel (see GET)"},
	{Text: "RAD", Description: "Switches trigonometric functions to compute in radians (see DEG)"},
	{Text: "READ", Description: "Reads data from a DATA statement"},
	{Text: "REM", Description: "Marks a comment in a program"},
	{Text: "RESTORE", Description: "Sets the position of where to read data from a DATA statement"},
	{Text: "RETURN", Description: "Ends a subroutine, effectively branching to the line immediately following the calling GOSUB (see GOSUB and POP)"},
	{Text: "RND", Description: "Returns a pseudorandom number"},
	{Text: "RUN", Description: "Starts execution of a program, optionally loading it from an external device (see LOAD)"},
	{Text: "SAVE", Description: "Writes a tokenized program to an external device; usually a cassette tape or disk (see LOAD)"},
	{Text: "SETCOLOR", Description: "Maps a logical color to a physical color"},
	{Text: "SGN", Description: "Returns the signum of a number"},
	{Text: "SIN", Description: "Returns the sine of a number"},
	{Text: "SOUND", Description: "Starts or stops playing a tone on a sound channel (see END)"},
	{Text: "SQR", Description: "Returns the square root of a number"},
	{Text: "STATUS", Description: "Returns the status of an I/O channel"},
	{Text: "STEP", Description: "Indicates the increment used in a FOR loop"},
	{Text: "STICK", Description: "Returns a joystick position"},
	{Text: "STOP", Description: "Stops the program, allowing later resumption (see CONT)"},
	{Text: "STRIG", Description: "Indicates whether a joystick trigger is pressed or not"},
	{Text: "STR", Description: "Converts a number to string form"},
	{Text: "THEN", Description: "Indicates the statements to execute if the condition is true in an IF statement"},
	{Text: "TO", Description: "Indicates the limiting condition in a FOR statement"},
	{Text: "TRAP", Description: "Sets to jump to a given program line if an error occurs (TRAP 40000 cancels this order)"},
	{Text: "USR", Description: "Calls a machine code routine, optionally with parameters"},
	{Text: "VAL", Description: "Returns the numeric value of a string"},
	{Text: "XIO", Description: "General-purpose I/O routine (from Fill screen to Rename file to Format disk instructions) "},
	{Text: "exit", Description: "Quit the application"},
	{Text: "quit", Description: "Quit the application"},
}

func executor(t string) {
	switch t {
	case "exit":
		fallthrough
	case "quit":
		os.Exit(0)
	default:
		println("Nothing happens")
	}
	return
}

func completer(in prompt.Document) []prompt.Suggest {
	if in.GetWordBeforeCursor() == "" {
		return nil
	} else {
		return prompt.FilterFuzzy(suggestions, in.GetWordBeforeCursor(), true)
	}
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
