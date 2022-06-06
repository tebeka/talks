error [done]
io.Reader/io.Writer
    - sha1 [done]
    - rot13 [done] (?)
    - io.Pipe

range request
    - ReaderAt
    - Seeker
    - io.Copy + io.LimitReader + io.Discard

fmt.Stringer [done]
fmt.Formatter [done]
json.Marshaler [done]

RoundTripper (mock, log) [done]
http.Flusher [done]
flag.Value [done]
testing.TB [done]

---
encoding.TextMarshaler
sort.Interface
heap.Interface
expvar.Var
sync.Locker
fs
http.Handler & http.HandlerFunc
