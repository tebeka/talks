## Python Optimization [½ Day Workshop]

Most of the time Python is fast enough. However in some cases you'll need to
make Python faster to meet business requirements.

In this workshop I'll teach tip & tricks from 20+ years experience of making
Python run faster. We'll covert the following topics:

- The Rules of Optimization Club
- Big O notation & why hardware matters
- Profiling & benchmarking
- Optimization tricks & common mistakes
- Using numba & Cython
- Integrating performance in your development process

## Open Sourcing Your Project [Full day workshop]

You wrote some cool code and would like to show off. You decide to open source
it, but - how?

In this workshop, we'll start with a small Python project (but feel free to
bring your own) and make it a full fledged open source project. We'll talk code
structure, picking a licence, testing & CI, versioning, packaging (setup.py) &
uploading to PyPI, documentation and more.

Most of the topics you'll learn at this workshop will be applicable to internal
projects at work as well.


## So You Think You Can Print?

You use Python's printing & string formatting almost on a daily basis, and most
of the time you don't give it much though. However you can do much more with
printing.

In this talk you'll learn the capabilities of string formatting (both
f-strings/.format and % formatting), see how to give give your objects a string
representation with `__str__` and `__repr__` (and why we have both). You'll
also see you you can use `__format__`, how to speed up logging and learn some
gotchas in printing.

## IPython - The Lost IDE

IPython seems like a fancy Python shell. Why do we need it when we have PyChar, VSCode, and other IDEs (IDE = integrated development environment).

I this talk you'll learn how to use the power of IPython for rapid development
and how you can integrate it with existing tools. We'll cover magic commands,
calling external process, usage of extended history and more. You'll also learn
on some popular extension and cool configuration hacks (such as `%autoreload
2`)

## Metaclasses and why You Shouldn't Use Them

Tim Peters said: "metaclasses are deeper magic than 99% of users should ever
worry about. If you wonder whether you need them, you don't."

In this talk I'll explain what are metaclasses and what are some possible uses
form them. We'll write our own implementation of `abc.ABC` using metaclass as
an example.

Lastly, I'll try to convince you why you shouldn't use metaclasses and point
out some other tools you might use instead (such as class decorators,
`__subclass_init__` and others)

At the end, if you'll think you're the 1% and end up using a metaclass - don't
say I didn't warn you! ☺
