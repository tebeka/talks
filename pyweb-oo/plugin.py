class LoggingPlugin:
    def shutdown(self):
        print('logger shutting down')


class SecurityPlugin:
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
