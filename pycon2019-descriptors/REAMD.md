# Descriptors

Talk on descriptors given at PyCon Israel 2019

## Script
- Live code?
    - Log terminal
    - %pycat
- Talk on `getattr`
- Show implementation without descriptors
- Show descriptors (binary file mode)
- docfile = '~/Projects/cpython/Doc/reference/datamodel.rst'
  !tail -n +1609 $docfile | less
- Show implementation with descriptors
- Start with `StaticMethod`, `ClassMethod`, `Property`?
- Show that methods are also descriptors
- Stock example 
    - Try to spot the bugs

## TODO
- References?
- Find other uses for descriptors
- ASCII first slide
-  %alias code pygmentize %l | cat -n                                                                                                                                                                                                       

