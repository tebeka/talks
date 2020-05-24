tag = 353/ep2020

build-docker:
	docker build -t $(tag) .

run-docker:
	docker run --rm -it -v ${PWD}:/ep2020 $(tag) /bin/bash
