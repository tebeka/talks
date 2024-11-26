%run setup.py

## Attribute access

%code player.py
%run player.py %code find_attr.py
%run find_attr.py
find_attr(p1, 'name')
find_attr(p1, 'count')
del p1.count
find_attr(p1, 'count')
%edit player.py (Player.count += 1)

instance knows only attributes + class. All the smart is in the class
mention properties, descriptors

## Slots
%code cars.py
%run cars.py
%edit cars.py (add `__slots__`)
%run cars.py
cars[0].license = 'PYTH0N'

## Mangling
%code cart.py
%run cart.py
%edit cart.py (`__`)
vars(cart)

## Static

%code vm.py
create VM for new id
global - need to import 2 things
staticmethod (rewrite attribute access) (`VM.new_id.__get__`)
Rule of thumb: not using self -> staticmethod

## Class

%code user.py
talk on adding json in ctor, bad design
method - need to create user first
statis - show admin
classmethod

!pydoc datetime.datetime (search for Class)

## ABC

%code plugin.py
%show cost.png
Add ABC


## Emulating Built-In Types

%code stack.py

!pydoc collections.abc
!pydoc collections.abc.Sequence

## Metaclasses

type([])
type(list)

in class:
- name
- bases
- mapping (methods, attributes)

%code point.py
%code robot.py

add 10 "def m1(self): pass" 

%fin
