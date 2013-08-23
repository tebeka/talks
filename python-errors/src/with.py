with open('/tmp/dummy.log', 'at') as out:  # <1>
    out.write('{}: ALIVE'.format(ctime()))

# The above equals to

out = None
try:
    out = open('/tmp/dummy.log', 'at')
    out.write('{}: ALIVE'.format(ctime()))
finally:
    if out:
        out.close()
