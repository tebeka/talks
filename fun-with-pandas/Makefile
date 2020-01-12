%.html: %.md
	kramdown -i GFM \
	    --template template.erb \
	    --syntax-highlighter rouge  $< > $@


mdignore=playbook.md
mdfiles=$(filter-out $(mdignore),$(wildcard *.md))
mdout=$(subst .md,.html,$(mdfiles))
class=cml
bucket=gs://353solutions/c

all: $(mdout) html

.PHONY: html
html:
	genhtml

.PHONY: clean
clean: clean-pyc
	rm $(mdout)

.PHONY: clean-pyc
clean-pyc:
	find . -name '*.pyc' -exec rm {} \;
	find . -depth -name __pycache__ -exec rm -rf {} \;

fresh: clean all

.PHONY: sync
sync: all
	@echo pushing to $(bucket)/$(class)
	@rsync \
	    --exclude .ipynb_checkpoints/ \
	    --exclude .git \
	    --exclude .gitignore \
	    --exclude Makefile \
	    --exclude REAMDE.md \
	    --exclude __pycache__ \
	    --exclude slides/ \
	    --exclude taxi.csv.bz2 \
	    --exclude template.erb \
	    -av . /tmp/$(class)
	@gsutil -m rsync -r /tmp/$(class) $(bucket)/$(class)
	@gsutil -q -m acl -r ch -u AllUsers:R $(bucket)/$(class)
