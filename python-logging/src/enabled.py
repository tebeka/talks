import logging
from time import sleep

def slow():
    sleep(100)
    return 7

log = logging.getLogger()

if log.isEnabledFor(logging.DEBUG):
    log.debug('slow() returned {}'.format(slow()))
