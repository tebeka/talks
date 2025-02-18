Session title

Looping in Go: More than just a "for"

Description

Quiz: What will the following print?

    ch := make(chan int)
    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
        }
    }()

    for v := range ch {
        fmt.Printf("%d ", v)
    }

Try it out at https://go.dev/play/p/of9h9opY_VS,, did you guess right?

In this talk we'll cover the various ways you can loop in Go and cover their semantics.
We'll loop over integers, slices, maps and channels and see the various semantics they present.
Finally, we'll cover the new looping directive that came in Go 1.22 and take a look at the good old `goto`.
No dinosaurs¹ will be harmed during this talk.

[1] https://xkcd.com/292/

---

Show loop assembly
