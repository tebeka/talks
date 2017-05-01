# Overview
* Why?
    * Isolated environment
	* Avoid DLL hell
    * Virtualization "heavy"
    * Standard way of shipping
	* Container metaphor
    * Written in Go
	* See https://www.youtube.com/watch?v=HPuvDm8IC-4
    * Go static exec - but we have assets, templates ...
* Based on LXC, layered FS, cgroups ...
    * FreeBSD jails
* images and containers
    * https://docs.docker.com/engine/docker-overview/#docker-engine
    * image - read only template (executable)
    * registry - holds images
	* https://hub.docker.com
    * container - running image (process)
* Suite of application - engine, compose, swarm, kubernetes ...
* Problems
    * https://www.youtube.com/watch?v=PivpCKEiQOQ
    * Orchestration
    * Size

# Run

* `docker version`
* `docker pull alpine`
* `docker run alpine`
    * error
* `docker run alpine hostname`
    * `hostname`
* `docker run alpine /bin/sh`
    * Nothing happens
* `docker run -it alpine /bin/sh`
    * Look around file system
    * `docker ps` in another window
    * `docker ps --format '{{.ID}} (from {{.Names}}) is running {{.Command}}'`
	* Go templates
    * `while true; do date && sleep 2; done`
    * `docker logs -ft $(docker ps -q)`
    * `sleep 100`
	- On host `ps aux | grep sleep`
* `docker ps -a`
    * See container
    * `docker rm`
    * `docker run --rm -it alpine /bin/sh`
* `docker pull python:alpine`
    * `docker run python:alpine python -m http.server`
	* Try hitting localhost:8000, no connection
	    * `docker exec -it $(docker ps -q) /bin/sh`
	    * `curl localhost:8000
    * `docker run --rm  -p 8000 python:alpine python -m http.server`
	* Try hitting localhost:8000, no connection
	* See port in `docker ps`
    * `docker run --rm -p 8080:8000 python:alpine python -m http.server`
* `docker run -it --rm -v ${PWD}:/code alpine /bin/sh`
    * `ls /code`
* `docker pull redis:alpine`
* `docker run -d --name redis redis:alpine`
* `docker run -it --rm redis:alpine /bin/sh`
    * `redis-cli`
	* can't connect
* `docker run --rm -it --link redis:db1 redis:alpine /bin/sh`
    * `redist-cli -h db1`
	* `GET hi`
	    * https://www.youtube.com/watch?v=119Gn8a6eiw&t=1m7s :)
	* `SET hi there`
	* `GET hi`
    * Exit, re-run and `GET hi`
* happy/Dockerfile
* `docker build -t happy .`
* `docker run happy`
    - fail
* `docker run --link 405aae4dcbfd:db1 -e HAPPY_REDIS=db1 happy`
* happy/docker-compose.yml
    * `docker-compose build`
    * `docker-compose up`
* We also have stack, swarm and kubernetes
