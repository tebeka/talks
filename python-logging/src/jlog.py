import logging, json

class JSONFormatter(logging.Formatter):
    def format(self, record):
        return json.dumps(vars(record))

handler = logging.StreamHandler()
handler.setFormatter(JSONFormatter())
log = logging.getLogger()
log.addHandler(handler)
log.fatal('Please reinstall universe and reboot')
