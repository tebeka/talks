#!/bin/bash

hg bookmark -r default master
hg push git+ssh://git@github.com/tebeka/nrsc.git
