def make_counter():
    x = [0]  # <1>
    def counter():
        x[0] += 1
        return x[0]
    return counter

c = make_counter()
print(c())
# => 1
print(c())
# => 2
