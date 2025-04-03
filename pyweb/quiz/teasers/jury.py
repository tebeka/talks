# Twelve Angry Man

from concurrent.futures import ProcessPoolExecutor

guilty = 0


def juror(id):
    global guilty

    print(f'J{id} ', end='')
    guilty += 1


with ProcessPoolExecutor() as pool:
    for i in range(12):
        pool.submit(juror, i)

print(guilty)
