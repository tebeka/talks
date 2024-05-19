# Square Peg in a Round Hole

from abc import abstractmethod, ABC


class Plugin(ABC):
    @abstractmethod
    def handle(self, msg):
        ...


class LoggingPlugin(Plugin):
    def handle(self, msg, id):
        print(f'{id}: {msg}')


lp = LoggingPlugin()
lp.handle('hello', 1)
