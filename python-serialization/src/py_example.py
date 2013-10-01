import config

assert config.points[0] == {'x': 0, 'y': 1}

cfg = {}
execfile('src/config.py', {}, cfg)  # <1>
assert cfg['points'] == config.points, 'corrupted data'
