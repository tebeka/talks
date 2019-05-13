from dataclasses import dataclass


@dataclass
class Proxy:
    name: str
    proxied: object

    def __getattr__(self, attr):
        if hasattr(self.proxied, attr):
            return getattr(self.proxied, attr)
        raise AttributeError(attr)
