#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    const char str[] = "This is a wonderful competition, is it not?  I really like ISTS.  It really brings out the best and the worst in people.  Look to your left and look to your right.  The people at your table are your team.  The people at other tables are your enemy.  There's a room full of world class Pentesters here itching to break your boxes.  You'll be lucky if don't they ask you to say \"Meow meow meow.  I'm a pretty kitty\" while you buy them things from the vending machine in exchange for your boxes back.  Have fun with this challenge, because at this point Red Team has probably burned your box to the ground.";
    char c;
    char * res;

    FILE * fp = fopen("flag.txt", "r");
    char flag[15];

    fscanf(fp, "%s", flag);

    for(size_t i = 0; i < strlen(flag); i++) {
        c = flag[i];
        res = (char *)malloc(strlen(str));

        for(size_t j = 0; j < (size_t)(strlen(str)); j++){
            res[j] = str[j] ^ c;
            printf("%x", res[j]);
        }

        printf("\n\n============================================================\n\n");
        free(res);
    }
}

