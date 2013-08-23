def format_message(type, value, traceback):
    message = StringIO()
    out = lambda m: message.write(u'{}\n'.format(m))

    out(ctime())
    out('== Traceback ==')
    out(''.join(format_exception(type, value, traceback)))  # <1>
    out('\n== Command line ==')
    out(' '.join(sys.argv))
    out('\n== Environment ==')
    for key, value in environ.items():
        out('{} = {}'.format(key, value))

    return message.getvalue()

