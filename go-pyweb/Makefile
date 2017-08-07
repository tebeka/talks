all: libsqrt.so

libsqrt.so: sqrt.go
	go build -buildmode=c-shared -o $@ $<

clean:
	rm libsqrt.*

present:
	 go get -u golang.org/x/tools/cmd/present
	 present
