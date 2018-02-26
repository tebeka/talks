# 4 Languages You Should Learn

PyWeb-IL March 2018 talk

---

Prolog (SWI)
Assembly (x86?, dosbox) - show qbart, num2str
Mainstream (Python, Go, C, C++, Java, Ruby ...)
Functional (Clojure, Haskell, SBCL, Racket)

## Notes
- Factorial in all languages?
- hw size (as vs gcc)
- speed?
- lisp 2 namespaces

    `(defun inc (n) (+ n 1))`
    `(defvar inc :hi)`

- `byte-seq` in clj-digest and `iter` on sockets 
- show influence on Python (Lang code, then python implementation)
    e.g. def fact(n): return reduce(mul, range(1, n+1))
- [Purely Functional Data Structures](https://www.cs.cmu.edu/~rwh/theses/okasaki.pdf)
- Assembly -> bit count problem (fixed list in memory)
- Framer crossing river in prolog
    - https://cseweb.ucsd.edu/classes/fa09/cse130/misc/prolog/goat_etc.html 
- sum tree (recursion)
    - tree.py
- Norvig code
    - https://github.com/norvig/paip-lisp/blob/master/lisp/prolog.lisp
    - https://github.com/norvig/paip-lisp
    - https://norvig.com/spell-correct.html
    - http://norvig.com/lispy.html
    - https://news.ycombinator.com/item?id=1803815
- https://github.com/kanaka/mal (make a lisp)
- Prolog 8 queens
    - https://gist.github.com/mfilej/136951
- Prolog
    - member(1, [1,2,3]).
    - member(X, [1,2,3]).
    - member(2, [1,X,3]).



# Books

* [Practical Common Lisp](http://www.gigamonkeys.com/book/)
* [On Lisp](http://www.paulgraham.com/onlisp.html)
* [Learn You a Haskell for Great Good](http://learnyouahaskell.com/chapters)
* [Real World Haskell](http://book.realworldhaskell.org/read/) 
* [SICP](https://mitpress.mit.edu/sicp/full-text/book/book-Z-H-4.html#%_toc_start)
    http://www.dabeaz.com/chicago/sicp.html
    https://twitter.com/dabeaz/status/948971981050966016
* [Prolog Experiments in Discrete Mathematics](http://samples.jbpub.com/9780763772062/PrologLabBook09.pdf) (PDF)
* [Assembly Programming Tutorial](https://www.tutorialspoint.com/assembly_programming/index.htm) (using [nasm](http://www.nasm.us/)
* [Programming from the Ground Up](http://nongnu.askapache.com/pgubook/ProgrammingGroundUp-1-0-booksize.pdf) (PDF)
* [OOC](https://www.cs.rit.edu/~ats/books/ooc.pdf) PDF

