all:
	$(error pick a target)


.PHONY: logs
logs:
	rm -rf logs
	mkdir -p logs/{1,2,3}
	# ftp://ita.ee.lbl.gov/traces/NASA_access_log_Aug95.gz
	zcat _data/NASA_access_log_Jul95.gz | split -l 100000 -d
	mv x0{0,1,2,3} logs/1
	mv x0{4,5,6,7} logs/2
	mv x* logs/3
	./_data/multiply.sh
	du -sh logs

query:
	curl -i --raw 'http://localhost:8080/logs?limit=10&filter=status%3E%3D400'
