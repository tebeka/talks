from typing import Protocol

class Plugin(Protocol):
    def shutdown(self):
        pass

class LoggingPlugin(Plugin):
    def shutdown(self):
        print('logger shutting down')


class SecurityPlugin(Plugin):
    def shutdown(self):
        print('security shutting down')


def shutdown(plugins: list[Plugin]):
    for plugin in plugins:
        plugin.shutdown(None)



class Event:
    def __init__(self, user, action):
        self.user = user
        self.action = action

plugins = [LoggingPlugin(), SecurityPlugin()]
