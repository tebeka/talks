import ctypes

_sqrt_dll = ctypes.cdll.LoadLibrary('./_sqrt.so')

sqrt = _sqrt_dll.sqrt
sqrt.argtypes = [ctypes.c_double]
sqrt.restype = ctypes.c_double

# Example usage
print(sqrt(2.0))
