#!/bin/bash

for i in {1..5}; do
	cp -r logs/1 logs/1-$i
	cp -r logs/2 logs/2-$i
	cp -r logs/3 logs/3-$i
done
