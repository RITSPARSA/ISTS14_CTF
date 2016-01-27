package main

import (
    "fmt"
    //"os"
    //"io"
    "io/ioutil"
    "strings"
    "bytes"
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

func shuffle(input []byte) []byte {
    s1 := input[0:5]
    s2 := input[5:10]
    s3 := input[10:15]
    s4 := input[15:20]
    
    //fmt.Println("1:", string(s1))
    //fmt.Println("2:", string(s2))
    //fmt.Println("3:", string(s3))
    //fmt.Println("4:", string(s4))

    var buffer bytes.Buffer

    buffer.Write(s3)
    buffer.Write(s1)
    buffer.Write(s4)
    buffer.Write(s2)

    return []byte(buffer.String())
}

func unShuffle(input []byte) []byte {
    s1 := input[0:5]
    s2 := input[5:10]
    s3 := input[10:15]
    s4 := input[15:20]
    
    //fmt.Println("1:", string(s1))
    //fmt.Println("2:", string(s2))
    //fmt.Println("3:", string(s3))
    //fmt.Println("4:", string(s4))

    var buffer bytes.Buffer

    buffer.Write(s2)
    buffer.Write(s4)
    buffer.Write(s1)
    buffer.Write(s3)

    return []byte(buffer.String())
}

func mangle(str string) string {
    message := []byte(str)
    message = shuffle(message)
    message = shiftLeft(2, message)
    return string(message)
    
}

func main() {
    //Read in ./flag.txt
    data, _ := ioutil.ReadFile("./flag.txt")
    data = []byte(strings.TrimSpace(string(data)))

    dataB64 := b64.StdEncoding.EncodeToString(data)
    
    for i := 0; i < 65; i++ {
        if len(dataB64) % 4 != 0 {
            fmt.Println("You can't do it")
            break
        }
        dataB64 = b64.StdEncoding.EncodeToString([]byte(dataB64))
        fmt.Printf("%d.) Len: %d\n", i, len(dataB64))
    }

    fmt.Println("Len:", len(dataB64))

    //shuffled := shuffle([]byte(dataB64))
    //fmt.Println(string(shuffled))

    //unShuffled := unShuffle(shuffled)
    //fmt.Println("Unshuffled:", string(unShuffled))
    //mangle(dataB64)
    /*fmt.Println(mangle(dataB64))
    fmt.Println("Placeholder")
*/
    /*for 1 - 1000 
        //Base64 Encode 
        //Mangle
    */
    //Write to file
    /*
    mangledDec, _ := b64.StdEncoding.DecodeString(mangled)
    fmt.Println(string(mangledDec))
*/
}
