# Setting Goals
- first performance goals
- percentile

# CPU

## Benchmarking & Profiling
    In []: %run outliers.py
    In []: data = pd.Series(np.random.randint(50, 60, 10_000))
    In []: data[7] = 3
    In []: data[1003] = 100
    In []: %timeit find_outliers(data)
    In []: %run nlp.py
    In []: s = 'We will encourage you to develop the three great virtues of a programmer: laziness, impatience, and hubris.'
    In []: %prun tokenize(s)
    In []: %%prun 
       ...: for _ in range(10_000): 
       ...:     tokenize(s) 
       ...:

https://wiki.c2.com/?RulesOfOptimizationClub
line_profiler
snakeviz

## Vectorization

    In []: import pandas as pd
    In []: s = pd.Series(range(10000))
    In [] : %%timeit 
       ...: total = 0 
       ...: for val in s: 
       ...:     total += val
       ...:  
    968 µs ± 35.9 µs per loop (mean ± std. dev. of 7 runs, 1000 loops each)
    In []: %timeit s.sum()
    49 µs ± 158 ns per loop (mean ± std. dev. of 7 runs, 10000 loops each)
    In []: 968 / 49
    Out[]: 19.755102040816325
    
## Boolean Indexing

    In []: import sqlite3
    In []: conn = sqlite3.connect('logs.db', detect_types=sqlite3.PARSE_DECLTYPES)
    In []: df = pd.read_sql('SELECT * FROM logs', conn)
    In []: len(df)
    Out[]: 10000

    In []: %%timeit 
	...: total = 0 
	...: for _, row in df.iterrows(): 
	...:     if row['status_code'] >= 400: 
	...:         total += 1 
	...:
    845 ms ± 3.23 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)
    In []: %timeit len(df[df['status_code'] >= 400])
    456 µs ± 1.14 µs per loop (mean ± std. dev. of 7 runs, 1000 loops each)
    In []: 845 / 0.456
    Out[]: 1853.0701754385964

## ufuncs

    In []: s = pd.Series(range(10000))
    In []: %timeit max(s)
    741 µs ± 10.5 µs per loop (mean ± std. dev. of 7 runs, 1000 loops each)
    In []: %timeit s.max()
    40.5 µs ± 128 ns per loop (mean ± std. dev. of 7 runs, 10000 loops each)
    In [34]: 741/40.5
    Out[34]: 18.296296296296298

## Appending

    In []: %run parse_log.py
    In []: import pandas as pd
    In [16]: %%timeit 
	...: df = pd.DataFrame() 
	...: for line in lines: 
	...:     df = df.append(parse_line(line), ignore_index=True) 
	...:
    2.65 s ± 5.73 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)
    In []: %timeit pd.DataFrame.from_records(parse_line(line) for line in lines)
    16 ms ± 1.1 ms per loop (mean ± std. dev. of 7 runs, 100 loops each)
    In []: 2650/17
    Out[]: 155.88235294117646

## Row iteration

    In []: import pandas as pd
    In []: import numpy as np
    In []: size = 50_000
    In []: df = pd.DataFrame({ 
	...:     'a': np.random.randint(1, 1000, size), 
	...:     'b': np.random.randint(1, 1000, size), 
	...:     'c': np.random.randint(1, 1000, size), 
	...: })   

The naive way will be to iterate over the rows, find the maximal value for each
row and add it to the total

    In []: total = 0
    In []: for _, row in df.iterrows(): 
	...:     total += row.max() 
	...:
    In []: total
    Out[]: 37413739

If you're following along, you'll see a different total since we use randomly
generated values.

Let's time it. I'm going to use the time magic and not the timeit magic since
this takes a while

    In []: %%time
	...: total = 0 
	...: for _, row in df.iterrows(): 
	...:     total += row.max() 

    CPU times: user 4.62 s, sys: 106 ms, total: 4.73 s
    Wall time: 4.63 s
    In []: %timeit df.apply(np.max, axis=1).sum()
    2.3 s ± 102 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)

    In []: %timeit df.apply(np.max, axis=1, raw=True).sum()
    194 ms ± 8.58 ms per loop (mean ± std. dev. of 7 runs, 10 loops each)

    In []: 4630/194
    Out[]: 23.8659793814433

## query

    In []: import sqlite3
    In []: conn = sqlite3.connect('logs2.db',
                detect_types=sqlite3.PARSE_DECLTYPES)
    In []: import pandas as pd
    In []: df = pd.concat([df] * 1000)
    In []: %timeit df[(df['method'] == 'GET') & (df['status_code'] >= 400]
    574 ms ± 5.42 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)
    In []: %timeit df.query('method == "GET" & status_code >= 400')
    336 ms ± 4.55 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)
    In []: 336/574
    Out[]: 0.58

## numba & Cython
    In []: def add(a, b):
       ...:     return a + b
       ...:
    In []: from dis import dis
    In []: dis(add)
      2           0 LOAD_FAST                0 (a)
		  2 LOAD_FAST                1 (b)
		  4 BINARY_ADD
		  6 RETURN_VALUE

    In []: import numba
    In []: @numba.jit
       ...: def jit_add(a, b):
       ...:     return a + b
       ...:

    In []: type(jit_add)
    Out[]: numba.core.registry.CPUDispatcher
    In []: jit_add.overloads
    Out[]: OrderedDict()
    In []: jit_add(1, 2)
    Out[]: 3
    In []: jit_add.overloads
    In []: jit_add.overloads.keys()
    Out[]: odict_keys([(int64, int64)])
    In []: f = jit_add.get_overload((numba.int64, numba.int64))
    In []: f
    Out[]: <function __main__._Closure.jit_add>
    In []: type(f)
    Out[]: builtin_function_or_method
    In []: f(1, 2)
    Out[]: 3
    In []: jit_add(1.0, 2.0)
    Out[]: 3.0
    In []: jit_add.overloads.keys()
    Out[]: odict_keys([(int64, int64), (float64, float64)])

    In []: import pandas as pd
    In []: import numpy as np
    In []: s = pd.Series(np.random.randint(-3, 200, 1_000_000))
    In []: def relu(n):
	...:     if n <= 0:
	...:         return 0
	...:     return n
	...:

    In []: %timeit s.apply(relu)
    294 ms ± 3.49 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)

    In []: @numba.vectorize
	...: def vect_relu(n):
	...:     if n <= 0:
	...:         return 0
	...:     return n
	...:

    In []: vect_relu(3)
    Out[]: 3

    In []: vect_relu(pd.Series([-1, 2, 3]))
    Out[]: array([0, 2, 3])

    In []: np.allclose(s.apply(relu), vect_relu(s))
    Out[]: True

    In []: %timeit vect_relu(s)
    946 µs ± 5.85 µs per loop (mean ± std. dev. of 7 runs, 1000 loops each)

    In []: 294000/946
    Out[]: 310.7822410147991

# Memory
- Why it's important?
    - Stay on one machine
    - swapping

## Measuring

    In []: mb = 1<<20
    In []: df = pd.read_csv('taxi.csv.xz')
    In []: df.dtypes
    In []: df.memory_usage()
    In []: df.memory_usage(deep=True)
    Out[]: 
    Index                         128
    VendorID                  3999992
    tpep_pickup_datetime     37999924
    tpep_dropoff_datetime    37999924
    passenger_count           3999992

    In []: df.memory_usage(deep=True).sum()/mb

## Loading Parts

    In []: df = pd.read_csv('taxi.csv.xz', usecols=['VendorID', 'total_amount'])
    In []: df.memory_usage(deep=True).sum()/mb

    In []: pd.read_csv('taxi.csv.xz', usecols=['VendorID', 'total_amount'],
        ...: chunksize=100_000)
    In []: for df in pd.read_csv('taxi.csv.xz', usecols=['VendorID',
	...:  'total_amount'], ...: chunksize=100_000):
	...:     print(len(df))

## Categorical
    In []: vendor_names = {
	...:     1: 'Creative',
	...:     2: 'VeriFone',
	...:     4: 'BigApple',
	...: }
    In []: vendors = df['VendorID'].map(vendor_names)
    In []: vendors.sample(5)
    In [93]: cat_vendors = vendors.astype('category')
    In [103]: cat_vendors.memory_usage(deep=True) /

Also check dtypes (float64 -> float32)
              vendors.memory_usage(deep=True)
    Out[103]: 0.015

Also Cython
    - %%cython

# Serialization
- SQL
- HDF5
- Parquet

# Alternate Data Frames
- dask
- Vaex 
