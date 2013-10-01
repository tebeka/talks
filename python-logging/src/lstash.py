import logging, logstash, math

log = logging.getLogger()
log.addHandler(logstash.LogstashHandler('localhost', 5959))

try:
    math.log(0)
except ValueError as err:
    log.exception('log(0) -> {}'.format(err))
