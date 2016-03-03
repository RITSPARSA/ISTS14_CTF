package main

import (
    "container/ring"
    "fmt"
    b64 "encoding/base64"
)

func main() {
    myArr := string([]rune{
        0x49, 0x53, 0x54, 0x53, 0x2d, 0x59, 0x41, 0x52,
        0x41, 0x2d, 0x45, 0x4e, 0x55, 0x52, 0x0a,
    })
    myStr := "Haha!!!!  Strings won't find this flag. You need to use something else!!"
    promiseRing := ring.New(15)
    
    result := make([]byte, len(myStr))
    restore := result
    
    fmt.Println(myStr)

    for i := 0; i < promiseRing.Len(); i++ {
        result = restore
        for j := 0; j < len(myStr); j++ {
            result[j] = myStr[j] ^ myArr[i]
        }
        promiseRing.Value = b64.StdEncoding.EncodeToString(result)
        promiseRing = promiseRing.Next()
    }

    fmt.Println("============")
    promiseRing.Do(func( x interface{}) {
        fmt.Print(x, "\n============\n")
    })

}
