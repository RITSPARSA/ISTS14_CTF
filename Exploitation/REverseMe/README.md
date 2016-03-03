#Let's Go Reversin'

##Overview
The files README-CTF.md and reverseMe.exe are the only files that will be provided to the competitors.  The binary is a 32-bit Windows PE compiled on Windows 7 with a source code langauge of C.  The code creates an integer array that contains the flag.  This means that when the program was compiled, it will not store it as a string in the executable.  This prevents competitors from finding it with the strings utility.  The code XORs the str[] variable with one byte from the flag at a time and prints that to the screen.  Since the input is not provided, they cannot find the original plaintext which is the flag.  This is to divert the competitor.  In reality, this flag can be found with a bit of searching though a hexdump of the data, or realizing that the program stores it as an integer array, and that it will store that differently.

##Flag
ISTS-INTA-RRAY
