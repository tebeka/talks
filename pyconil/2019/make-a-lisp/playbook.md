- log
- ipython
- font
- silence phone
- slides up to `Code`

~~~
    $ python scm.py
    > (define inc (lambda (n) (+ n 1)))
    > (inc 10)
~~~

- Code
    - Who said code need to be written? (scratch, labview) - 2d
    - Who said code need to be in files? (pharo)
    - Python 3 UTF-8
- Lex
    - Convert bytes to tokens
    - %code tok.py
    ~~~
    code = b'x = collatz(n)`
    print_tokens(code)
    ~~~
    - show it ignore whitespace
    - %edit -x -n 1115 ~/Projects/cpython/Parser/tokenizer.c
    - humble.py
    - scm.py:tokenize
	- token is just a str (file, lineno ...)
    ~~~
    code = '(+ 7 3)'
    tokenize(code)
    code = '(even? 2)'
    tokenize(code)
    ~~~
- Parse
    - %edit -x ~/Projects/cpython/Parser/Python.asdl
    - %edit -x ~/Projects/cpython/Grammar/Grammar
    - %code parse.py
    ~~~
    code = '''
    if x > 10:
        y /= 7
    '''
    - reader
    - scm.py:read_sexpr
    - All numbers are float (like JavaScript)
    ~~~
    code = '(* (+ 7 3) 12)'
    read_sexpr(tokenize(code))
    code = '(* 3 4) (+ 5 6)'
    tokens = tokenize(code)
    read_sexpr(tokens)
    read_sexpr(tokens)
    ~~~
- Eval
    - Python `eval('8 * 9')
    - story on assember
    - expression vs statement `eval('x=1')` → `exec('x=1')`
    - start with just func, name & variable
	- need builtin
	- names + case sensitive
	- lisp 1/2
    - Python's `__add__` and friends
    - run & repl
	- Python str & repr
	- readline for matching ()
    - if
	- what are booleans? (we'll use 1.0 & 0.0)
	- short circuit
	- in Python as well
~~~
	if True:
	    'OK'
	else:
	    1/0

	def myif(cond, true_val, false_val):
	    if cond:
		return true_val
	    return false_val
~~~
    - or, and
	- they do it
	- `(or)` -> 0.0
	- `(and)` -> 1.0
	- `True or 1/0`
    - define & set!
    - lambda
	- parameter passing (value, ref ..)
	- square.scm
	- collatz.scm
	- fact.scm (recursion)
	- adder.scm (closure)
	    - Show Python
    - account.scm (global, nonlocal)
	- Python, start without nonlocal
    - begin
