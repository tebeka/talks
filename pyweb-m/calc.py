"""Calculate math expressions"""
import math


def calc(expr):
    """Returns the calculation of a math expression"""
    env = globals().copy()
    env.update(math.__dict__)
    return eval(expr, env)


def main():
    from argparse import ArgumentParser

    parser = ArgumentParser(description=__doc__)
    parser.add_argument('expr', help='expression to calculate')
    args = parser.parse_args()

    try:
        print(calc(args.expr))
    except Exception as err:
        raise SystemExit(f'error: {args.expr!r} - {err}')


if __name__ == '__main__':
    main()
