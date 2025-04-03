"""Calculate math expressions"""
import math


def calc(expr):
    """Returns the calculation of a math expression"""
    return eval(expr, math.__dict__)


def main():
    from argparse import ArgumentParser

    parser = ArgumentParser(description=__doc__)
    parser.add_argument('expr', help='expression to calculate')
    args = parser.parse_args()

    try:
        val = calc(args.expr)
        print(val)
    except Exception as err:
        raise SystemExit(f'error: {args.expr!r} - {err}')


if __name__ == '__main__':
    main()
