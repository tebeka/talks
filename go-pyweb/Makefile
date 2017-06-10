%.so: %.go
	go build -buildmode=c-shared -o $@ $<

all: fetch.so gmath.so
