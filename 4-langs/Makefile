dotfile = /tmp/graph.dot

graph.png: graph.pl
	echo 'digraph G {' > $(dotfile)
	echo 'graph [rankdir=LR, label="Sample Graph", labelloc=t, fontsize=32];' >> $(dotfile)
	echo 'edge [fontsize=10, color=blue];' >> $(dotfile)
	sed \
	    -e 's/edge(\(\w\),\(\w\))\./\1 -> \2;/' \
	    -e s'/path.*//' < $< >> $(dotfile)
	echo '}' >> $(dotfile)
	dot -Tpng -o$@ < $(dotfile)

hw: hw.o
	ld $< -o $@

hw.o: hw.s
	as -o $@ $<


clean:
	rm -f graph.png
	rm -f hw.o hw
