def load_config(filename):
    with open(filename) as fp:
        code = fp.read()
    cfg = {}
    exec(code, None, cfg)
    return cfg
