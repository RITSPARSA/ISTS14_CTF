#include <stdio.h>
#include <stdlib.h>
#include <string.h>


unsigned int arr[] = {
    0x49, 0x53, 0x54, 0x53, 0x2d, 0x49, 0x4e, 0x54, 0x41,
    0x2d, 0x52, 0x52, 0x41, 0x59, 0x0a
};

int main() {
    const unsigned int arrLen = 15;
    const char str[] = "Haha!!!!  You thought you could find the flag for ISTS CTF with strings.";
    char c;
    char * res;

    printf("%s\n", str);

    for(size_t i = 0; i < arrLen; i++) {
        c = arr[i];
        res = (char *)malloc(strlen(str));

        for(size_t j = 0; j < (size_t)(strlen(str)); j++){
            res[j] = str[j] ^ c;
        }
        printf("%s\n", res);
        free(res);
    }
}

