#TextMangler
Go program that mangles some text to reveal a flag.

##Directories
There are 2 directories for this challenge.  The 'prod' directory is the directory that is given to competitors.  The 'solution' directory is houses the solution to the challenge, as well as the initial PoC code.

###Prod
The 'prod' directory is what will be given to competitors to complete.  The competitors are given three files:

1. README.md
    - README.md outlines the basics and small amount of background on the challenge.
2. TextMangler.go
    - TextMangler.go is a Go source file containing obfuscated code that reads a string from a file, mangles it, and then writes it to a file.
3. output.txt
    - output.txt is the result of mangling the flag string.


###Solution
This is a solution directory.  This directory contains three files:

1. TextMangler.go
    - TextMangler.go is a Go source file containing the original source code that reads a string from a file, mangles it, and then writes it to a file.  This also contains the functions that that unmangle the string, but does not employ them.  This was the original code with comments.
2. fixIt.go
    - fixIt.go is a Go source file that reads in the output of TextMangler.go, unmangles the string, and then prints it to a file called "flag.txt".  This i an example solution written in Go.  The solution may be in any language, as long as the flag is obtained.
3. flag.txt
    - flag.txt is the flag that results from the output of fixIt.go.  This is the flag that will result in points once submitted.

Please extract prod/output.txt.tar.gz and use that as input for fixIt.go.

##Writeup
The competitors are given three files:

1. README.md
    - README.md outlines the basics and small amount of background on the challenge.
2. TextMangler.go
    - TextMangler.go is a Go source file containing obfuscated code that reads a string from a file, mangles it, and then writes it to a file.
3. output.txt
    - output.txt is the result of mangling the flag string.

The competitor has to "reverse engineer", decipher, de-obfuscate, or whatever you want to call it the Go source code in TextMangler.go.  The competitor should realize that the flag is being read in, having all excess white space trimmed from the end, mangled, and then written to a file.

When viewing the function call to mangle(), the competitor will see that there is a loop that runs 65 times.  In that loop, the message is base64 encoded, rotated left by 5 spaces, and then has a shuffle function performed on it.

The shuffle function divides the message into 4 equal parts s1, s2, s3, s4 and then places them in the order s3, s1, s4, s2.

All the competitor has to do, is perform the operations of TextMangler in reverse the same number of times to obtain the flag.  The solution can be seen in solution/fixIt.go.
