import module
print(module.submodule.value)
# => 1
# Now change submodule value to 7
reload(module)
print(module.submodule.value)
# => 1
