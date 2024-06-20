#!/bin/bash
#
ag --vimgrep -s --go 'goto\s+' ~/sdk/go1.22.0/src | \
    grep -v testdata  | \
    grep -v _test.go | \
    wc -l
