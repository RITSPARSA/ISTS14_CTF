#Let's Go Reversin'

##Overview
The files README-CTF.md and GoReversin.exe are the only files that will be provided to the competitors.  The binary is a 32-bit Windows PE compiled on Windows 7 with a source code langauge of Go.  The code creates a rune array that contains the flag.  This means that when the program was compiled, it will not store it as a string even though it is cast as a string in the program.  This prevents competitors from finding it with the strings utility.  The code creates a ring data type, or a circularly joined doubly linked list, and XORs the myStr variable with each one of the bytes in the key and prints it to the screen after base64 encoding it.  Since the input is not provided, they cannot find the original plaintext which is the flag.  This is to divert the competitor.  In reality, this flag can be found with a bit of searching though a hexdump of the data, or realizing that the program is compiled from Go and will store strings differently and searching according to those customs.

##Flag
ISTS-YARA-ENUR

