def load_config(filename):
    """Load configuration from Python file"""
    with open(filename) as fp:
        code = fp.read()
    cfg = {}
    exec(code, None, cfg)
    return cfg
