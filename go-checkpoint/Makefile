html/%.html: code/%.go
	pygmentize -Ofull,linenos=1 -l go -f html -o $@ $<


gofiles=$(wildcard code/*.go)
goout=$(subst .go,.html,$(subst code/,html/,$(gofiles)))

all: $(goout)

clean:
	rm -f html/*.html

fresh: clean all
