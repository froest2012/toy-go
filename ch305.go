package main

import (
    "fmt"
    "unicode/utf8"
    "strings"
    "bytes"
    "strconv"
)

func main() {
    s := "hello, world"
    fmt.Println(len(s))
    fmt.Println(string(s[0]), string(s[7]) + "a")
    c := s[len(s) - 1]
    fmt.Println(c)
    fmt.Println(s[0:5])
    fmt.Println(s[:5])
    fmt.Println(s[7:])
    fmt.Println(s[:])
    fmt.Println("goodbye" + s[5:])
    s = "left foot"
    t := s
    s += ", right foot"
    fmt.Println(t, s)

    const GoUsage = `Go is a tool for managing Go source code.

    Usage:
        go command [arguments]
    ...`
    fmt.Println(GoUsage)
    fmt.Println(HasPrefix(s, "hello"))
    fmt.Println(HasSuffix(s, "world"))
    fmt.Println(HasSuffix(s, "foot"))
    fmt.Println(Contains(s, "foot"))
    s = "hello, 世界"
    fmt.Println(len(s))
    fmt.Println(utf8.RuneCountInString(s))
    for j := 0; j < len(s); {
        r, size := utf8.DecodeRuneInString(s[j:])
        fmt.Printf("%d\t%c\n", j, r)
        j += size
    }
    fmt.Println(len(s[7:]))
    fmt.Println(utf8.DecodeRuneInString(s[7:]))
    for m, n := range "Hello, 世界" {
        fmt.Printf("%d\t%q\t%d\n", m, n, n)
    }
    fmt.Println(string(0x4eac))
    fmt.Println(string(65))
    fmt.Printf("%x\n", string(1234567))
    fmt.Println(basename("a/b/c.go"))
    fmt.Println(basename("a.b.c.go"))
    fmt.Println(basename("abc"))
    fmt.Println(basename1("a/b/c.go"))
    fmt.Println(basename1("a.b.c.go"))
    fmt.Println(basename1("abc"))
    fmt.Println(comma("18394774598938747774885"))
    s = "abc"
    b := []byte(s)
    s2 := string(b)
    fmt.Println(b)
    b[1] = 33
    fmt.Println(s2)
    fmt.Println(b)
    values := []int{21,33,121,48}
    fmt.Println(intsToString(values))
    fmt.Println(comma1("1837494575889378757"))
    x := 123
    y := fmt.Sprintf("I'm string %d", x)
    fmt.Println(y, strconv.Itoa(x))
    fmt.Println(strconv.FormatInt(int64(x), 2))
}

func HasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) &&  s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

func Contains(s, substr string) bool {
    for i := 0; i < len(s); i++ {
        if HasPrefix(s[i:], substr) {
            return true
        }
    }
    return false
}

func basename(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    return s
}

func basename1(s string) string {
    slash := strings.LastIndex(s, "/")
    s = s[slash + 1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(",")
        }
        fmt.Fprintf(&buf, "%c", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

func comma1(s string) string {
    var buf bytes.Buffer
    n1 := len(s)%3
    buf.WriteString(s[:n1])
    for i := n1; i < len(s); i += 3 {
        buf.WriteString(",")
        fmt.Fprintf(&buf, "%s", s[i:i + 3])
    }
    return buf.String()
}
