# Conventions in Go

Uppercase exported
Testing
    `_test.go`
    TestXXX
    BechmarkXXX
    ExampleXXX
    FuzzXXX
    TestMain
    testdata
build
    `_GOOS.go`
Doc comment
    // Sqrt returns the square root of v
    func Sqrt(v float64) float64 { ... }
iota in const group
    const (
        Reader Role = iota + 1
        Writer
        Admin
    )
"internal" package

https://twitter.com/tebeka/status/1609825350510608385
