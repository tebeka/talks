%.so: %.go
	go build -buildmode=c-shared -o $@ $<

all: fetch.so gmath.so

present:
	 go get -u golang.org/x/tools/cmd/present
	 present
