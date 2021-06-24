count:
	@echo -n 'count: '
	@grep .play quiz-codebar.slide | grep -v '^#' | wc -l
	@# @grep .play quiz-gcil.slide | grep -v '^#' | wc -l
	@# @grep .play quiz.slide | grep -v '^#' | wc -l

run:
	present 

test:
	go test -v slide_test.go
