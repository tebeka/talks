# Getting Ready
:ALEDisable

# Setup
- outliers.go:NewOutliers
- glue.c:load_func, talk on refcount (py_last_error)
- outliers.go:Close

# Down
- outliers.go:Detect (up to line 67 - KeepAlive)
- glue.c:detect

# Up
- outliers.go:Detect (from line 78 - last error)
- cArrToSlice design & ownership

# Building
- outliers.go:11
- Makefile
