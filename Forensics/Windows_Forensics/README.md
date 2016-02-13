### Memory Forensics 1
* This is the .vmem file of a WindowsSP0 machine.
* Volatility however will report it as Windows7SP0/SP1 either should be accepted.

[Memory Image 1 Location](https://drive.google.com/file/d/0B8Tamy_f__H9S3JHZ0RMN25IdjA/view?usp=sharing)

===
####  Questions
* Q) What is this a memory image of? (100 points)
    * A) Windows7SP0
```
    volatility -f Image_Forensics_1.vmem imageinfo
```

* Q) What is the IP address of the machine? (100 points)
    * A) 10.80.100.31
    * Various ways to accomplish, one way listed below.

```
    volatility -f Image_Forensics_1.vmem --profile Windows7SP0 consoles
```

* Q) What program has PID of 3520? (100 points)
    * A) putty.exe
```
    volatility -f Image_Forensics_1.vmem --profile=Windows7SP0 pslist
```


