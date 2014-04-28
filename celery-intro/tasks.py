from celery import Celery
from celery.utils.log import get_task_logger
from celery.contrib import rdb

from time import sleep


celery = Celery(
    broker='amqp://guest@localhost',  # Task queue
    backend='amqp',  # Results
)
log = get_task_logger(__name__)


@celery.task  # Task decorator
def add(x, y):
    sleep(1)  # Simulate work
    log.info('adding {} to {}'.format(x, y))
    return x + y


@celery.task  # Task decorator
def div(x, y):
    sleep(1)  # Simulate work
    log.info('divigin {} by {}'.format(x, y))
    return x / y


@celery.task  # Task decorator
def mul(x, y):
    sleep(1)  # Simulate work
    log.info('multiplying {} to {}'.format(x, y))
    rdb.set_trace()
    return x * y
