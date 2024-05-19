#!/bin/bash

# for flf in /usr/share/figlet/fonts/*.flf; do
for flf in /usr/share/figlet/*.tlf; do
    font=$(basename $flf)
    font=${font/.tlf/}
    echo $font
    echo IPython | toilet -f $font
done
