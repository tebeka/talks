count:
	@echo -n 'count: '
	@grep -F .play quiz-gcuk.slide | grep -Fv // | wc -l
	@# @grep .play quiz-codebar.slide | grep -v '^#' | wc -l
	@# @grep .play quiz-gcil.slide | grep -v '^#' | wc -l
	@# @grep .play quiz.slide | grep -v '^#' | wc -l

run:
	present 

test:
	go test -v slide_test.go

slide-files:
	cat quiz-gcuk.slide | awk '/.play/ { ($$1 == "//") ? f=$$5 : f=$$4; print f }'
