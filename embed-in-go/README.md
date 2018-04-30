# Embedding Other Languages in Go

Talk given at [Go Israel meetup][meetup].

[meetup]: https://www.meetup.com/Go-Israel/events/kjvczlyxhbdb/


## Points
- What is nuclio
- Go
    - inject code
	- dependency problem
    - plugin
	- nuclio.Event != nuclio.Event
- Python
    - JSON/HTTP
    - One process per worker (GIL)
    - Unix socket
- pypy
    - RPC runtime
    - cffi
    - stack traces
- nodejs
    - failed embed (v8 worked)
    - JSON/HTTP  
- Java
    - JSON/HTTP
    - No unix socket
    - jackson slow
    - CLASSPATH to load user handler problem
    - embed in jar
- Future
    - Python
	- Embed
	- GIL
	- Sub interpreters
	- Load DLL 
    - Common FFI runtime
    - Better RPC
	- encoding (flatbuffers, msgpack ...)
	- shared memory

- Benchmakrs?
