ag --vimgrep -s --go 'goto\s+' ~/sdk/go1.19.5/src | \
    grep -v testdata  | \
    grep -v _test.go | \
    wc -l