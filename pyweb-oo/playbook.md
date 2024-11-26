%code player.py
%run player.py
%code find_attr.py
%run find_attr.py
find_attr(p1, 'name')
find_attr(p1, 'count')
del p1.count
find_attr(p1, 'count')
%edit player.py (Player.count += 1)
%code
%code cars.py
%run cars.py
%edit cars.py (add `__slots__`)
%run cars.py
cars[0].license = 'PYTH0N'
%code cart.py
%run cart.py
%edit cart.py (`__`)
vars(cart)
