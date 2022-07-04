from setuptools import setup

setup(
    name='calc',
    version='0.1.0',
    description='Command line calculator',
    author='Miki Tebeka',
    author_email='miki@353solutions.com',
    py_modules=['calc'],
    entry_points={
        'console_scripts': ['calc=calc:main'],
    },
)
