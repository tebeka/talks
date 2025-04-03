# Meta

# Who Go in PyWeb
- If all you have is a hammer
    - If all you have is C++ :)
- Programming language is a tool, not religion
- Side note: Try
    - Main stream (Python, Ruby, Perl, C, C++, Java ...)
    - Assembly
    - Functional (Lisp, ML, Haskell ...)
    - Logic (Prolog)

# A Personal Journey
- The free lunch - march 2005
    - Customer ask for speed, you wait for intel
    - stopped, now cores
- C10K
    - thread pool
    - async (Twisted asyncio), headache

# Looked at
- Erlang: Ericson, telephony, OTP, light weight process, mailbox, supervisor, ...
- Clojure: Lisp, JVM, STM, core.async as library
- Go: C Based, simple, goroutines, channels

# Go
- 2007 in Google. Ken Thompson (C), Rob Pike (plan 9), Robert Griesemer (v8) 
- Went for the C++, got Python + Ruby
- Rob Pike story on C++ header includes thousands
- Linux, Mac and Windows(?). x86, x86 and ARM
- go and GCC front end

# Why Go
- Static
    - Iron.io from 30 RoR servers to 2 Go
- Fast compilation
    - stdlib compiles in 0.3 seconds on my machine
- GC: Design simplicity


# Hello World
- package
- import
- func
- unicode
- ;

# Hello World 2
- types after name
- `*http.Request` - have pointers, not pointer semantic
    - like C++ reference
    - unsafe
- Fprint: exports
- net/http -> http.ListenAndServe

# Declaration
- const
    - iota
- new(Point) -> 0 initializtion
- When running says there's code we don't see on slide

# Syntax
- anonymous functions
- type at end

# range
- single and double context
    - line enumerate

# Defer and Errors
- How would you do that with exceptions?
- Defered are LIFO
- Error and last return value
- More verbose but more robust
- err in if context
- Change name

# Goroutines

# Channels

# Example
- Comment out close - show deadlock

# Select
- Change timeout

# Method
- Talk about pointers vs copy

# Interfaces
- Many in stdlib (io.Writer)
- interface{} -> void *

# sha256
- file implements io.Reader
- hash implements io.Writer

# A TCP Proxy
- Write this in Python for scale :)


# Packages

# The go Tool
- GOPATH =~ virtualenv
- test simple but good enough
- fetch: coming dep tool

# Interact with Python [1]
- Go is used a lot as internal services

# Interact with Python [2]
- Build a DLL

# Interact with Python [3]
- Run "make"
    - Show we got `libsqrt.*`
- Run sqrt.py from shell

# Interact with Python [4]


<!--
vim: ft=markdown spell
-->
