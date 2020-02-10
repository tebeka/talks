## Python Optimization [½ Day Workshop]

Most of the time Python is fast enough. However in some cases you'll need your
code to run faster in order to meet business requirements.

In this workshop you'll learn tips & tricks from my 20+ years experience of
making Python run faster. Some of the topics we'll cover are:

- The Rules of Optimization Club
- Big O notation & why hardware matters
- Profiling & benchmarking
- Optimization tricks & common mistakes
- Using numba & Cython
- Integrating performance in your development process

## Open Sourcing Your Project [Full day workshop]

You wrote some cool code and decided to open source it so others can enjoy it.
But... how?

In this workshop, we'll start with a small Python project (feel free to bring
your own instead) and make it a full fledged open source project. We'll talk
code structure, picking a licence, testing & CI, versioning, packaging
(setup.py) & uploading to PyPI, documentation and more.

Most of the topics you'll learn at this workshop will be applicable to internal
projects at work as well.

## So You Think You Can Print?

You use Python's printing & string formatting almost on a daily basis, most of
the time without giving it much thought. However there's much more to it.

In this talk you'll learn the capabilities of string formatting (both
f-strings/.format and % formatting), see how to give your objects a string
representation with `__str__` and `__repr__` (and why both exist). We'll
also cover `__format__`, how to speed up logging and learn about some gotchas
in printing.

## IPython: The Productivity Booster

IPython seems like a fancy Python shell. Why do we need it when we have
PyCharm, VSCode, and other IDEs[1]?

In this talk you'll learn how to use the power of IPython for rapid development
and how you can integrate it with existing tools. We'll cover magic commands,
calling external process, usage of extended history, async/await and more.
You'll also learn on some popular extension and cool configuration hacks (such
as `%autoreload 2`)

Since Jupyter is based on IPython, you'll be able to use all of what you
learned in Jupyter Lab/Notebooks as well.

[1]: IDE = integrated development environment

## Metaclasses and why You Shouldn't Use Them

Tim Peters said: "metaclasses are deeper magic than 99% of users should ever
worry about. If you wonder whether you need them, you don't."

In this talk I'll explain what are metaclasses and what are some possible uses
for them. We'll write our own implementation of `abc.ABC` using metaclass as an
example.

Lastly, I'll convince you why you shouldn't use metaclasses and point out some
other tools you might use instead (such as class decorators,
`__subclass_init__` and others)

At the end, if you'll still think you're the 1% and end up using a metaclass -
don't say I didn't warn you! ☺
