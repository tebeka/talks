import logging as log

log.basicConfig(
    filename='/tmp/basic.log',  # <1>
    format='[%(asctime)s] %(levelname)s %(filename)s:%(lineno)s %(message)s',
    level=log.INFO,
)

log.info('How are you?')

# Output (in /tmp/basic.log)
# [2013-09-30 16:02:58,927] INFO basiccfg.py:9 How are you?
