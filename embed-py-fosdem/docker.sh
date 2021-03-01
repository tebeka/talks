#!/bin/bash

tag=fosdem2021

case $1 in
    build ) docker build -t ${tag} . ;;
    run ) docker run --rm -it -v ${PWD}:/home/fosdem/outliers ${tag};;
    * ) echo "usage: $(basename $0) build|run"; exit 1;;
esac
