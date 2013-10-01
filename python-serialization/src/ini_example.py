from ConfigParser import ConfigParser

cfg = ConfigParser()
section, age = 'daffy', 76
cfg.add_section(section)
cfg.set(section, 'type', 'duck')
cfg.set(section, 'age', age)

tmp = '/tmp/looney.ini'
with open(tmp, 'w') as out:
    cfg.write(out)

cfg = ConfigParser()
with open(tmp) as fo:
    cfg.readfp(fo)

assert cfg.getint(section, 'age') == age, 'corrupted data'  # <1>
