from setuptools import setup
import re


def read_version():
    with open('math2/__init__.py') as fp:
        for line in fp:
            match = re.search(r"__version__\s*=\s*'([^']+)'", line)
            if match:
                return match.group(1)


setup(
    name='math2',
    version=read_version(),
    author='Miki Tebeka',
    author_email='miki@353solutions.com',
    url='https://github.com/tebeka/talks/tree/master/vim-rocks/math2',
    packages=['math2'],
)
