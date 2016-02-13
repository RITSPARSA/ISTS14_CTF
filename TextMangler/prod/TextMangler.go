package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "bytes"
    b "encoding/base64"
)

func g(n int, a []byte) []byte {
    m := a
    var t byte

    for i := 0; i < n; i++ {
        for j := 0; j < len(m) ; j++ {
            if j == 0 {
                t = m[0]
            }
            
            if j == len(m) - 1 {
                m[j] = t
            } else {
                m[j], m[j + 1] = m[j + 1], m[j]
            }
        }
    }
    return m
}

func d(w []byte) []byte {
    h:= len(w)
    y := h/ 2
    e := y / 2
    k := y + e

    l := w[:e]
    z := w[e:y]
    n := w[y:k]
    u := w[k:h]
    
    var q bytes.Buffer

    q.Write(n)
    q.Write(l)
    q.Write(u)
    q.Write(z)

    return []byte(q.String())
}

func m(i string) string {
    f := []byte(i)
    for i := 0; i < 65; i++ {
        fmt.Printf("%d.) Len: %d\n", i, len(f))
        f = []byte(b.StdEncoding.EncodeToString(f))
        f = g(5, f)
        f = d(f)
    }
    return string(f)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, _ := ioutil.ReadFile("./flag.txt")
    str := strings.TrimSpace(string(data))
    
    f, err := os.Create("output.txt")
    check(err)

    defer f.Close()

    mangled := m(str)
    _, e := f.WriteString(mangled)
    check(e)
}
