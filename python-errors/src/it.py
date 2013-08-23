def ints(upto):
    for i in xrange(upto):
        if i == 3:
            cont = (yield -1)  # Get value back
            yield None  # Caller will discard
            if not cont:
                return
        yield i

to10 = ints(10)
for i in to10:
    if i == -1:
        to10.send(True)  # Signal to continue, ignore return value
        continue
    print(i)

