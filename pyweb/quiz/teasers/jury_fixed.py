# Twelve Angry Man

from concurrent.futures import ProcessPoolExecutor
from multiprocessing import Value, Lock

lock = Lock()
guilty = Value('i', 0)


def juror(id):
    print(f'J{id} ', end='')
    with lock:
        guilty.value += 1


with ProcessPoolExecutor() as pool:
    for i in range(12):
        pool.submit(juror, i)

print(guilty.value)
