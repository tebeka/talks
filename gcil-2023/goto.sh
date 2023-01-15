ag --vimgrep -s --go 'goto\s+' ~/sdk/go/src | \
    grep -v testdata  | \
    grep -v _test.go | \
    wc -l
