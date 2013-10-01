import logging

log = logging.getLogger('looney')
log.setLevel(logging.INFO)
handler = logging.FileHandler('/tmp/hard.log')
log.addHandler(handler)

fmt = '[%(asctime)s] %(levelname)s %(filename)s:%(lineno)s %(message)s'
formatter = logging.Formatter(fmt)
handler.setFormatter(formatter)

log.info('How are you?')

# Output (in /tmp/hard.log)
# [2013-09-30 16:39:24,810] INFO hard_way.py:12 How are you?
