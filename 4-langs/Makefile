dotfile = /tmp/graph.dot
s3bucket=s3://353solutions/4-langs

%: %.s
	as -o $@.o $<
	ld -o $@ $@.o


# white?
slides.html: slides.md
	pandoc -t revealjs -V theme=serif -s $< -o $@

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

fact: fact.s

clean:
	rm -f graph.png
	rm -f hw.o hw
	rm -f fact.o fact

# TODO: Move to GitHub pages
upload:
	s3cmd sync *.png *.jpg slides.html reveal.js $(s3bucket)/
	s3cmd setacl $(s3bucket)/ --acl-public --recursive
	find reveal.js/ \
	    -name '*.css' \
	    -exec s3cmd modify $(s3bucket)/{} \
	    	--add-header='Content-Type: text/css' \;
