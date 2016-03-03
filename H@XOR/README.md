#H@XOR

##Overview
The files that shall be given to competitors shall be README-CTF.md, H@XOR.exe, and output.txt.

##Writeup
H@XOR.exe is a 32-bit binary compiled from C on a Windows 7 system.  The code is rather simple.  It reads in a flag from flag.txt and uses each sequential byte as a key for XORing a long string that is stored in the binary file.  One byte of the key is used per every line of text printed to the screen.  With some basic static analysis, the large string can be found, which is the initial plaintext.  Advanced static analysis of the binary shows that the Boolean XOR operation is being used to jumble the bits that are printed out in hexadecimal.  Since the competitor is given the plaintext and the cipertext, they are able to deduce the key that was used to XOR each line.  Appending all of these characters together gets you the flag.

##Boolean XOR Explanation
The following operation represents the situation where both the plaintext and the ciphertext is known, but the key is not.  In this case the key is the flag.

         0   0   1   1
    XOR  ?   ?   ?   ?
         -------------
         0   1   1   0

This is what it would look like if we were using the Boolean OR operation.

         0   0   1   1
    OR   ?   ?   ?   ?
         -------------
         0   1   1   1

Notice, how if we try to derive the key with the OR operation, the 8s column and the 4s columns are unambiguous.  We know since the input value is 0 a key value of 0 would yield 0 output and a key value of 1 would yield a 1 output.


         0   0   1   1
    OR   0   1   ?   ?
         -------------
         0   1   1   1

We run into a situation with the input values of 1, because only one 1 is necessary to yield a 1.  This prevents us from knowing whether the rightmost values are 1s or 0s due to the nature of the OR operation.  A similar problem is encountered with the AND operation, but with 0s instead of 1s.

         0   0   1   1
    OR   0   1  1|0 1|0
         -------------
         0   1   1   1


         0   0   1   1
    AND 1|0 1|0 1|0  1 
         -------------
         0   0   0   1

Since XOR requires exactly one 1 to yield a 1 output, this problem is not found with XOR.

         0   0   1   1
    XOR  0   1   0   1
         -------------
         0   1   1   0


