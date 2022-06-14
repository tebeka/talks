# Moonraker
data = '{"id":"","login":"M","created":"2955-04-05T00:00:00+00:00"}'
request = json.loads(data)
m = User(**request)
print('m:', m)
