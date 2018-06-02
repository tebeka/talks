#!/bin/bash
"""Convert video (mp4) to animated gif"""

case $1 in
    -h | --help ) echo "usage: $(basename $0) IN_FILE OUT_FILE"; exit;;
esac

if [ $# -ne 2 ]; then
    1>&2 echo "error: wrong number of arguments"
    exit 1
fi

tmp=$(mktemp -d)
ffmpeg -i $1 -r 5 "${tmp}/frame-%03d.jpg"
convert -delay 20 -loop 0 "${tmp}/frame-0*" $2
