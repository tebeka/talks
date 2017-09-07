#!/usr/bin/env python

import ctypes

_sqrt_dll = ctypes.cdll.LoadLibrary('./libsqrt.so')

# From generated "libsqrt.h"
sqrt = _sqrt_dll.sqrt
sqrt.argtypes = [ctypes.c_double]
sqrt.restype = ctypes.c_double

# Example usage
print(sqrt(2.0))
