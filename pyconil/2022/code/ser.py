def add(a, b):
    return a + b


val = add(1, 2)

# ----

import json


# client
request = {
    'fn': 'add',
    'args': (1, 2),
}
data = json.dumps(request)
send_to_server(data)

# server
request = json.loads(data)
fn = resolve(request['fn'])
out = fn(*request['args'])
data = json.dumps(out)
send_to_client(data)

# client
val = json.loads(data)


