## Python Optimization (Workshop)
- Baseline & requirements
- Explain O(n)
- timeit
- pytest-benchmark
- profiling
    - snakeviz
- tricks
    - local/global
    - search in list/set
    - deque
    - `[0]*size`
- tracemalloc
- numba
- Cython

## Open source a project (Workshop)
- github
- code structure
    - readme
- license
- testing
- versioning
- CI
    - banner in readme
- readthedocs
- packaging
    - setup.py, manifest.in
    - upload to pypi
    - dependabot

## Faster Pandas (Workshop)

- vectorization
- ufuncs
- append
- eval & query
- join(?)
- sql & hdf5
- numba & cython

## OO Workshop
- class & instance attributes
- getattr
- properties
- staticmethod
- classmethod
- `__slots__`

## So you think you can print?
- `__repr__` & `__str__`
- `__format__`
- `{name:<20}`
- variable width for the above
- logging & `%(name)s`
- don't `print` in libraries

## IPython for Faster Development
- %magic
- files = !ls, ls $logs
- history
- configuration (aggressive reload)
- extensions: sql, cython, line/memory profiles, write your own (%nuclio)

## Metaclasses and why you shouldn't use them
- What are they
- Implement ABC
