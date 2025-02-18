## Title

What's in a name? Understanding variable scope.

## Abstract (250 characters)

When you see `logging.info('user: %r', user)`, do you know where `user` comes from?
We'll dive into Python's scope rules, discuss mutation vs re-binding, `global`, `nonlocal` and more.

It might seem academic, but it'll save you from future bugs.


## Description

Python's scoping rules seem clear: local, closure, global, builtin (LCGB).
However, knowing where a specific variable comes from can be tricky.

We'll start by discussing what is a variable in Python and see the `globals` and `locals` dicts and the `__closure__` function attribute.
Then we'll move to LCGB, and see where the `global` and `nonlocal` keyword come to play.
Finally, we'll see some common scope related bugs and how to avoid them.
