from dis import dis


def relu(n):
    if n < 0.0:
        return 0.0
    return n


print(dis(relu))
