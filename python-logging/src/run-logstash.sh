#!/bin/bash
# You'll need to download the logstash jar somewhere
# http://logstash.net/

case $1 in
    -h | --help ) echo "usage: $(basename $0) [LOGSTASH_JAR]"; exit;;
esac

jar=${1-logstash-1.2.1-flatjar.jar}
java -jar logstash-1.2.1-flatjar.jar agent -f src/logstash.conf
