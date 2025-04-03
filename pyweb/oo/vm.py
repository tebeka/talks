from enum import Enum
from uuid import uuid4


class Status(Enum):
    INIT = 'init'
    STARTING = 'starting'
    RUNNING = 'running'
    STOPPED = 'stopped'


class VM:
    def __init__(self, image, cpu_count, mem_size):
        self.image = image
        self.cpu_count = cpu_count
        self.mem_size = mem_size
        self.status = Status.INIT

    def start(self):
        self.id = self.new_id()
        self.status = Status.STARTING

    @staticmethod
    def new_id():
        return uuid4().hex
