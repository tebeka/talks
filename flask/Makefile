# For this to work, you'll need asiidoc and pygments installed
# Also have the following in ~/.asciidoc/asciidoc.conf
#    [attributes]
#    pygments=

txt = flask.txt
html = $(subst .txt,.html,$(txt))


all: $(html)

$(html): $(txt) ${PWD}/style.css *.png
	asciidoc -a stylesheet=${PWD}/style.css $<

clean:
	rm -f $(html)
	find . -name db.sqlite -exec rm -v {} \;

fresh: clean all

.PHONY: all clean fresh
