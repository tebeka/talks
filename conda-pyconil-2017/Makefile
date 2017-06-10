all:
	$(error please pick a target)

docker:
	docker build -t pyconil .

docker-ub:
	docker build -t pyconil-ub -f Dockerfile.ubuntu .

docker-ub-pre:
	docker build -t pyconil-ub -f Dockerfile.ubuntu.pre-installed .
