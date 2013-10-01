#!/bin/bash
# You'll need to download and unpack kibana
# http://www.elasticsearch.org/overview/kibana/installation/


case $1 in
    -h | --help ) echo "usage: $(basename $0) [KIBANA_DIR]"; exit;;
esac

root=${1-.}
cd ${root}
python -m SimpleHTTPServer
