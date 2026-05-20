# Getting Ready

- Tab 1
    vim
    :ALEDisable
    set nospell

- Tab 2
    ./docker.sh run
    make  # first build is slow
    clear

# Setup
- show python code, example_test.go
- outliers.go:NewOutliers
- glue.c:load_func, talk on refcount (py_last_error), show Close
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
