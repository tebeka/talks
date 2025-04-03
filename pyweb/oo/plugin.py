from abc import ABC, abstractmethod

class Plugin(ABC):
    @abstractmethod
    def shutdown(self):
        ... # We'll never get here

class LoggingPlugin(Plugin):
    def shutdown(self):
        print('logger shutting down')


class SecurityPlugin(Plugin):
    def shutdwon(self):
        print('security shutting down')


def shutdown(plugins):
    for plugin in plugins:
        plugin.shutdown()



class Event:
    def __init__(self, user, action):
        self.user = user
        self.action = action

plugins = [LoggingPlugin(), SecurityPlugin()]
