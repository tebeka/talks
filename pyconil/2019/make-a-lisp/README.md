# Let's Write a Lisp Interpreter

Workshop at PyCon Israel 2019

## Pitch

[Greenspun's tenth
rule](https://en.wikipedia.org/wiki/Greenspun%27s_tenth_rule) states that `Any
sufficiently complicated C or Fortran program contains an ad-hoc,
informally-specified, bug-ridden, slow implementation of half of Common Lisp.`.
Understanding how programming language work will make you a better programmer
and gain a better understanding of Python itself.

We'll implement a small [lisp like](http://norvig.com/lispy.html) language and
discuss language design & implementation  issues and how they are found in
Python.
- [Lexing](https://en.wikipedia.org/wiki/Lexical_analysis) & Parsing: What are the implication of Python using whitespace for indentation?
- Variable scope & [closures](https://en.wikipedia.org/wiki/Closure_(computer_programming)): Why we have `global` and `nonlocal` in Python
- Types: Why the value of `1/2` changed from Python 2 to 3
- Evaluating code: Python's `eval` vs `exec` and byte code interpreter. Why does `or` and `and` [short curcuit](https://en.wikipedia.org/wiki/Short-circuit_evaluation)
