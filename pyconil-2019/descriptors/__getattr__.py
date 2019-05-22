class Proxy:
    def __init__(self, name, proxied):
        self.name = name
        self.proxied = proxied

    def __getattr__(self, attr):
        return getattr(self.proxied, attr)
