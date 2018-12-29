def sqrt(val, epsilon=0.00001, niter=10_000):
    """Return the square root of val using Newton's method"""
    if val < 0:
        raise ValueError('sqrt of negative number')

    if val == 0:
        return 0  # shortcut

    guess = 1.0
    for i in range(niter):
        if abs(guess*guess - val) <= epsilon:
            return guess
        guess = (val/guess + guess) / 2.0

    raise ValueError(f'cannot find a solution after {niter} iterations')
