package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "bytes"
    b64 "encoding/base64"
)

func rotateLeft(numIter int, arr []byte) []byte {
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

func rotateRight(numIter int, arr []byte) []byte {
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
    l1 := len(input)
    l2 := l1 / 2
    l3 := l2 / 2
    l4 := l2 + l3

    s1 := input[:l3]
    s2 := input[l3:l2]
    s3 := input[l2:l4]
    s4 := input[l4:l1]
    
    var buffer bytes.Buffer

    buffer.Write(s3)
    buffer.Write(s1)
    buffer.Write(s4)
    buffer.Write(s2)

    return []byte(buffer.String())
}

func unShuffle(input []byte) []byte {
    l1 := len(input)
    l2 := l1 / 2
    l3 := l2 / 2
    l4 := l2 + l3

    s1 := input[:l3]
    s2 := input[l3:l2]
    s3 := input[l2:l4]
    s4 := input[l4:l1]
    
    var buffer bytes.Buffer

    buffer.Write(s2)
    buffer.Write(s4)
    buffer.Write(s1)
    buffer.Write(s3)

    return []byte(buffer.String())
}

func mangle(str string) string {
    message := []byte(str)
    for i := 0; i < 65; i++ {
        fmt.Printf("%d.) Len: %d\n", i, len(message))
        message = []byte(b64.StdEncoding.EncodeToString(message))
        message = rotateLeft(5, message)
        message = shuffle(message)
    }
    return string(message)
}

func unMangle(str string) string {
    fmt.Println("Unmangling the string.")
    message := []byte(str)
    for i := 65; i > 0; i-- {
        fmt.Printf("%d.) Len: %d\n", i, len(message))
        message = unShuffle(message)
        message = rotateRight(5, message)
        message, _ = b64.StdEncoding.DecodeString(string(message))
    }
    return string(message)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    //Read in ./flag.txt, trim the spaces, and convert to string
    data, _ := ioutil.ReadFile("./flag.txt")
    str := strings.TrimSpace(string(data))
    
    f, err := os.Create("output.txt")
    check(err)

    defer f.Close()

    //Mangle the string
    mangled := mangle(str)
    _, e := f.WriteString(mangled)
    check(e)
}
