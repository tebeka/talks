import json

msg = (1, 2, 3)
data = json.dumps(msg)
reply = json.loads(data)
print(msg == reply)
