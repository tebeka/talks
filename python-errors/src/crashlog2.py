def install(emails=None, logfile=None):
    global _EMAILS, _PREV_EXCEPTHOOK, _LOGFILE

    _EMAILS = emails or _EMAILS
    _LOGFILE = logfile
    _PREV_EXCEPTHOOK = sys.excepthook
    sys.excepthook = excepthook
