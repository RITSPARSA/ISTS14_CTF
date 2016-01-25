package main

import (
    "fmt"
    //"os"
    b64 "encoding/base64"
)

func shiftLeft(numIter int, arr []byte) []byte {
    message := arr
    var temp byte

    for i := 0; i < numIter; i++ {
        for j := 0; j < len(message) ; j++ {
            if j == 0 {
                temp = message[0]
            }
            
            if j == len(message) - 1 {
                message[j] = temp
            } else {
                message[j], message[j + 1] = message[j + 1], message[j]
            }
        }
    }
    return message
}

func shiftRight(numIter int, arr []byte) []byte {
    message := arr
    var temp byte

    for i := 0; i < numIter; i++ {
        for j := len(message) - 1; j > 0 ; j-- {
            if j == len(message) - 1 {
                temp = message[len(message) - 1]
            }
            
            if j == 0 {
                message[j] = temp
            } else {
                message[j], message[j - 1] = message[j - 1], message[j]
            }
        }
    }
    return message
}

func mangle(str string) string {
    message := []byte(str)
    message = shiftLeft(2, message)

    return string(message)
    
}

func main() {
    var first string = "first"
    //var second string = "second"

    firstEnc := b64.StdEncoding.EncodeToString([]byte(first))
    fmt.Println(firstEnc)
    mangled := mangle(firstEnc)
    fmt.Println(mangled)
    /*
    mangledDec, _ := b64.StdEncoding.DecodeString(mangled)
    fmt.Println(string(mangledDec))
*/
}
