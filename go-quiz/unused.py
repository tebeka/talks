#!/usr/bin/env python

import sys
from pathlib import Path

# .play -edit -numbers stridx.go
used = {
    line.split()[-1]
    for line in sys.stdin
    if '.play' in line and '//' not in line
}

all = set(p.name for p in Path('.').glob('*.go'))
print('\n'.join(sorted(all - used)))
