call_count = 0


def handler(message):
    call_count += 1
    print(f'got {message}')
