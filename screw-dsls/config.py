port = 8080
db_host = 'db1.353solutions.com'


def _load_rc():
    from os import path, environ

    default = path.expanduser('~/.config/353/config')
    cfg_file = environ.get('CONFIG_FILE', default)
    if path.isfile(cfg_file):
        with open(cfg_file) as fp:
            exec(fp.read(), globals())


_load_rc()
del _load_rc
