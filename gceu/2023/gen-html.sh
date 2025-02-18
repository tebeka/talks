#!/bin/bash

mkdir -p html
for src in *.go; do
    cat $src | \
	sed -e 's# // HL##' -e 's/\t/    /g' | \
	grep -F -v OMIT | \
        pygmentize -Ofull,linenos=1,style=vs -l go -f html -o html/${src/.go/.html}
done
