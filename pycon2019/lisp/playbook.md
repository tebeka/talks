- Slides up to `Code`
- Code
    - Who said code need to be written? (scratch, labview) - 2d
    - Who said code need to be in files? (pharo)
    - Python 3 UTF-8
- Lex
    - Convert bytes to tokens
    - tok.py, `x = collatz(y)`
	- show it ignore whitespace
    - %edit -x -n 1145 ~/Projects/cpython/Parser/tokenizer.c:
    - scm.py:tokenize
	- token is just a str (file, lineno ...)
    ~~~
    code = '(+ 7 3)'
    tokenize(code)
    ~~~
- Parse
    - %edit -x ~/Projects/cpython/Parser/Python.asdl
    - parse.py
    - reader
    - scm.py:read_sexpr
    - All numbers are float (like JavaScript)
    ~~~
    read_sexpr(tokenize(code))
    code = '(* (+ 7 3) 12)'
    read_sexpr(tokenize(code))
    code = '(* 3 4) (+ 5 6)'
    tokens = tokenize(code)
    read_sexpr(code)
    read_sexpr(code)
    ~~~
- Eval
    - Python `eval('8 * 9')
    - story on assember
    - expression vs statement `eval('x=1')` → `exec('x=1')`
    - start with just func, name & variable
	- need builtin
    - run & repl
    - if
    - or, and
	- short circuit
	- in Python as well
    - define & set
    - lambda
	- square.scm
	- collatz.scm
	- adder.scm (closure)
	    - Show Python
	- fact.scm (recursion)
    - begin
	- account.scm (global, nonlocal)
