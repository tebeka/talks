'''Encode/decode numbers from/to base62.'''

__version__ = '0.1.1'

import string

chars = string.digits + string.ascii_letters
base = len(chars)


def encode(num):
    '''Encode number in base62, returns a string.'''
    if num < 0:
        raise ValueError('cannot encode negative numbers')

    if num == 0:
        return chars[0]

    digits = []
    while num:
        rem = num % base
        num = num // base
        digits.append(chars[rem])
    return ''.join(reversed(digits))


def decode(string):
    '''Decode a base62 string to a number.'''
    loc = chars.index
    size = len(string)
    num = 0

    for i, ch in enumerate(string, 1):
        num += loc(ch) * (base ** (size - i))

    return num


def main(argv=None):
    import sys
    from argparse import ArgumentParser

    argv = argv or sys.argv

    parser = ArgumentParser(description='encode/decode in base62')
    parser.add_argument('n', help='number to encode/decode')
    parser.add_argument('--encode', '-e', help='encode (default)',
                        action='store_true', default=False)
    parser.add_argument('--decode', '-d', help='decode', action='store_true',
                        default=False)
    args = parser.parse_args(argv[1:])

    if args.encode and args.decode:
        raise SystemExit('error: cannot have both -e and -d')

    if args.decode:
        print(decode(args.n))
        raise SystemExit

    try:
        n = int(args.n)
    except ValueError:
        raise SystemExit('error: bad number')

    print(encode(n))

if __name__ == '__main__':
    main()
