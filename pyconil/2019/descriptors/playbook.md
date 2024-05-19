# Prepare

- log
- ipython
- font
- slides
- silence phone

# Script

- Love it after lunch ...
- Myself + Tenerife
- We're going to dive into one of Python powerful tool that not many know about
- It helps explain how many features in Python, such as properties, static
  methods, class methods and others works
- We'll also see how it can help us easily implement features that otherwise
  are more difficult to write


    In [1]: fp = open('/dev/zero')

    In [4]: fp.mode
    Out[4]: 'r'

    In [2]: fp.__dict__
    Out[2]: {'mode': 'r'}

    In [3]: vars(fp)
    Out[3]: {'mode': 'r'}

    In [5]: dir(fp)

    In [5]:fp.read(3)
    Out[6]: '\x00\x00\x00'

    In [6]:fp.read
    Out[6]: <function TextIOWrapper.read(size=-1, /)>

    In [7]: fp.__class__.read
    Out[7]: <method 'read' of '_io.TextIOWrapper' objects>

    In [8]: %code getattr.py
    In [9]: get_attr(fp, 'mode')
    found in object
    Out[9]: 'r'

    In [10]: get_attr(fp, 'read')
    found in class TextIOWrapper
    Out[10]: <method 'read' of '_io.TextIOWrapper' objects>

    In [19]: %code __getattr__.py
    In [20]: p = Proxy('dev0', fp)

    In [21]: p
    Out[21]: Proxy(name='dev0', proxied=<_io.TextIOWrapper name='/dev/zero' mode='r' encoding='UTF-8'>)

    In [22]: p.name
    Out[22]: 'dev0'

    In [24]: p.read(3)
    Out[24]: '\x00\x00\x00'

    # There's also __getattribute__, Rachel!

    In [27]: fp.name
    Out[27]: '/dev/zero'

    In [28]: vars(fp)
    Out[28]: {'mode': 'r'}

    In [29]: fp.__class__.name
    Out[29]: <attribute 'name' of '_io.TextIOWrapper' objects>

    In [30]: dir(fp.__class__.name)
    In [31]: %edit -n 1609 -x ~/Projects/cpython/Doc/reference/datamodel.rst

    In [34]: %code prop_demo.py

    In [35]: p = Person('Diana', 'Prince')

    In [ ]: %show diana.jpg

    In [36]: p
    Out[36]: Person(first='Diana', last='Prince')

    In [37]: p.name
    Out[37]: 'Diana Prince'

    In [38]: %code property.py

    In [35]: p = PersonD('Natasha', 'Romanova')

    In [ ]: %show natasha.jpg

    In [37]: p.name
    Out[37]: 'Natasha Romanova'

    In [66]: PersonD.name
    Out[66]: <__main__.Property at 0x7fd157af7780>

    In [ ]: %code cls_demo.py
    In [30]: p = Point.from_json('{"x": 1, "y": 2}')
    In [31]: p.x
    Out[31]: 1

    In [32]: %code classmethod.py
    In [33]: p = PointD.from_json('{"x": 1, "y": 2}')

    In [34]: p.x
    Out[34]: 1

    In [35]: type(p)
    Out[35]: __main__.PointD

    In [40]: %code static_demo.py
    In [41]: Math.neg(10)
    Out[41]: -10

    In [42]:%code static_method.py
    In [41]: MathD.neg(10)
    Out[41]: -10

    In [  ]: from json import JSONDecoder
    In [  ]: '__get__' in dir(JSONDecoder.decode)
    In [  ]: # This is how "self" gets to methods

    In [  ]: %mb

    In [  ]: %cow WIIFM
    In [  ]: from trade import Trade
    In [55]: t = Trade('BRK.A', 310132.41, 1)
    In [  ]: t
    Out[22]: Trade('BRK-A', 300304.13, 1)
    In [57]: t.symbol
    Out[57]: 'BRK.A'
    In [57]: t.price
    Out[57]: 300304.13
    In [57]: t.volume
    Out[57]: 1
    In [59]: t.symbol = 'brka'
    In [60]: t.price = 'too much'
    In [  ]: t.volume = 2
    In [  ]: t.volume
    Out[  ]: 2
    In [58]: vars(t)
    Out[58]: {'_Trade_symbol': 'BRK.A', '_Trade_price': 310132.41, '_Trade_volume': 3}

    In [53]:%code trade.py
    In [54]: Trade.symbol
    Out[54]: <fields.Symbol at 0x7fd157a8d198>


    In [50]: %code fields.py
    In [47]: %code field.py (talk on ...)

    In [49]:f = Field()  # Abstract class


