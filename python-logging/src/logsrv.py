#!/usr/bin/env python

import logging as log
import logging.config as logcfg

port = 9999

if __name__ == '__main__':
    logcfg.fileConfig('src/log.conf')
    thr = logcfg.listen(port)
    thr.daemon = True  # Don't hangup
    thr.start()

    while True:
        line = raw_input('>>> ')
        level, message = line.split(' ', 1)
        fn = getattr(log, level, None)
        fn(message)
