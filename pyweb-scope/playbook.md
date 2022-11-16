%code geo.py
    - local: lat1, lng1, lat2, lng2, dlag, dlng
    - global: deg_km
    - builtin: abs


    distance.__code__.co_nlocals

%edit geo.py
    print(locals())

before return

distance(1, 2, 3, 4)

    distance.__code__.co_names

%edit -x geo.py
    print([k for k in globals() if k not in vars(__builtins__)])

!python geo.py

    from dis import dis
    dis(distance)


%code version.py
%code app.py

!python app.py

%code handler.py

    handler('login')

    from datetime import datetime
    call_times = []

    ...
    call_times.append(datetime.now())

add `global`

%code metrics.py

    handler('login')

add nonlocal

    handler.__code__.co_freevars
    handler.__closure__

%code for.py

    size_of('Python')
    size_of('')

%code click.py

    lambda i=i: print(i)

    def make_handler(n):
        def fn():
            print(n)
        return fn

    clicks = [make_handler(i) for i in range(10)]

%code comp.py

    n

%code ex.py

div(1, 0)


%edit ex.py
    err = None

div(1, 0)

%edit ex.py

    err = None
    excpet Exception as e:
        err = e

div(1, 0)

%code ann.py
    
    host


- # Objects are merely a poor man's closures. - Qc Na
%code bank.py
