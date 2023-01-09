ag --vimgrep -s --go goto ~/sdk/go1.19.4/src | \
    grep -v testdata  | \
    grep -v _test.go | \
    wc -l
