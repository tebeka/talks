from signal import signal, SIGTERM, SIGINT
import atexit

def bye():
    print('BYE')

atexit.register(bye)

def signal_handler(signum, stackframe):
    raise SystemExit

for signum in (SIGTERM, SIGINT):
    signal(signum, signal_handler)
