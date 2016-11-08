#!/bin/bash

bzip2 -d -c edw.sqlite.bz2 > edw.sqlite

virtualenv venv
. ./venv/bin/activate
pip install -r requirements.txt
