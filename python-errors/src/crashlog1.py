def excepthook(type, value, traceback):
    try:
        if not (_EMAILS or _LOGFILE):
            return

        message = format_message(type, value, traceback)
        if _EMAILS:
            send_email(_EMAILS, sys.argv[0], message)

        if _LOGFILE:
            with open(_LOGFILE, 'at') as fo:
                fo.write('{}\n'.format(message))

    finally:
        if _PREV_EXCEPTHOOK:
            _PREV_EXCEPTHOOK(type, value, traceback)
