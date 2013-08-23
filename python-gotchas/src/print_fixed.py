from __future__ import print_function # or use python3  <1>
verbose = 1
if verbose:
    log = lambda message: print(message)
else:
    log = lambda message: 1

# OR
if verbose:
    def log(message): print(message) # <2>
else:
    def log(message): pass
