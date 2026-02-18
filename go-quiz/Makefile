slide = quiz-gosyd.slide
count:
	@echo -n 'count: '
	@grep -F .play $(slide) | grep -Fv // | wc -l

run:
	present 

test:
	go test -v slide_test.go

slide-files:
	cat $(slide) | awk '/.play/ { ($$1 == "//") ? f=$$5 : f=$$4; print f }'

unused:
	./unused.py < $(slide)
