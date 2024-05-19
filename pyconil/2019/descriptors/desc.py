#!/usr/bin/env python
# coding: utf-8

# # Descriptors and the Python dot operator
# 
# Miki Tebeka, [353Solutions](http://353solutions.com)
# 
# * https://docs.python.org/3/reference/datamodel.html#object.__getattr__
# * https://docs.python.org/3/reference/datamodel.html#implementing-descriptors
# * https://docs.python.org/3/howto/descriptor.html

# In[75]:


fp = open('/dev/random')
fp


# In[76]:


fp.mode


# In[77]:


getattr(fp, 'mode')


# In[78]:


fp.__getattribute__('mode')


# In[79]:


fp.__dict__


# In[80]:


vars(fp)


# In[81]:


'mode' in vars(fp)


# In[85]:


fp.read(10)


# In[90]:


fp = open('/dev/random', 'rb')
fp


# In[91]:


fp.read(10)


# In[92]:


fp.mode


# In[93]:


'mode' in vars(fp)


# In[95]:


'mode' in vars(fp.__class__)


# In[103]:


fp.readlines


# In[104]:


'readlines' in vars(fp)


# In[105]:


'readlines' in vars(fp.__class__)


# In[106]:


fp.__class__.__bases__


# In[108]:


'readlines' in vars(fp.__class__.__bases__[0])


# In[109]:


fp.__class__.__mro__


# In[111]:


for cls in fp.__class__.__mro__:
    if 'readlines' in vars(cls):
        print(cls)
        break
else:
    print('no idea where "readlines" comes from')


# ## Descriptors

# In[120]:


class Desc:
    def __get__(self, inst, owner):
        print('__get__: inst=%r, owner=%r' % (inst, owner))
        
    def __set__(self, inst, value):
        print('__set__: inst=%r, value=%r' % (inst, value))
        

class Stock:
    symbol = Desc()
    
s = Stock()


# In[121]:


s.symbol


# In[122]:


Stock.symbol


# In[123]:


s.symbol = 'brk.a'


# In[125]:


Stock.symbol = 'brk.a'
Stock.symbol


# In[126]:


s.symbol


# In[1]:


class Field:
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return self._value
        
    def __set__(self, inst, value):
        self.assert_valid(value)
        self._value = value
        
    def assert_valid(self, value):  # Mention abc
        pass


class SymbolField(Field):
    def assert_valid(self, value):
        if not str.isupper(value):
            raise ValueError('symbol must be upper case, got %r' % value)
            

class PriceField(Field):
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError('price must be float, got %s' % type(value))
            
        if value <= 0:
            raise ValueError('price most be positive, got %s' % price)
        
    
class Stock:
    symbol = SymbolField()
    price = PriceField()
    
    def __init__(self, symbol, price):
        self.symbol = symbol
        self.price = price
        
    def __repr__(self):
        cls = self.__class__.__name__
        return '%s(%r, %r)' % (cls, self.symbol, self.price)

brka = Stock('BRK.A', 216298.00)
brka


# In[3]:


v = Stock('V', 97.48)
v.price


# In[4]:


brka.price


# In[140]:


class Field:
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return getattr(inst, '_field')
        
    def __set__(self, inst, value):
        self.assert_valid(value)
        setattr(inst, '_field', value)        
        
    def assert_valid(self, value):
        pass


class SymbolField(Field):
    def assert_valid(self, value):
        if not str.isupper(value):
            raise ValueError('symbol must be upper case, got %r' % value)
            

class PriceField(Field):
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError('price must be float, got %s' % type(value))
            
        if value <= 0:
            raise ValueError('price most be positive, got %s' % price)
        
    
class Stock:
    symbol = SymbolField()
    price = PriceField()
    
    def __init__(self, symbol, price):
        self.symbol = symbol
        self.price = price
        
    def __repr__(self):
        cls = self.__class__.__name__
        return '%s(%r, %r)' % (cls, self.symbol, self.price)

brka = Stock('BRK.A', 216298.00)
brka


# In[165]:


class Field:
    _attr = None
    
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return getattr(inst, self._attr)
        
    def __set__(self, inst, value):
        self.assert_valid(value)
        setattr(inst, self._attr, value)        
        
    def assert_valid(self, value):
        pass


class SymbolField(Field):
    _attr = '_symbol'
    
    def assert_valid(self, value):
        if not str.isupper(value):
            raise ValueError('symbol must be upper case, got %r' % value)
            

class PriceField(Field):
    _attr = '_price'
    
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError('price must be float, got %s' % type(value))
            
        if value <= 0:
            raise ValueError('price most be positive, got %s' % price)
        
    
class Stock:
    symbol = SymbolField()
    price = PriceField()
    
    def __init__(self, symbol, price):
        self.symbol = symbol
        self.price = price
        
    def __repr__(self):
        cls = self.__class__.__name__
        return '%s(%r, %r)' % (cls, self.symbol, self.price)

brka = Stock('BRK.A', 216298.00)
brka


# In[166]:


vars(brka)


# In[168]:


Stock('v', 74.48)


# In[148]:


class Field:
    _attr = None
    
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return getattr(inst, self._attr)
        
    def __set__(self, inst, value):
        self.assert_valid(value)
        setattr(inst, self._attr, value)        
        
    def assert_valid(self, value):
        pass


class SymbolField(Field):
    _attr = '_symbol'
    
    def assert_valid(self, value):
        if not str.isupper(value):
            raise ValueError('symbol must be upper case, got %r' % value)
            

class PriceField(Field):
    _attr = '_price'
    
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError('price must be float, got %s' % type(value))
            
        if value <= 0:
            raise ValueError('price most be positive, got %s' % price)
        
    
class Stock:
    symbol = SymbolField()
    price = PriceField()
    low = PriceField()
    
    def __init__(self, symbol, price, low):
        self.symbol = symbol
        self.price = price
        self.low = low
        
    def __repr__(self):
        cls = self.__class__.__name__
        return '%s(%r, %r, %r)' % (cls, self.symbol, self.price, self.low)

brka = Stock('BRK.A', 216298.00, 216297.00)
brka


# In[159]:


from itertools import count
class Field:
    _next_id = count().__next__

    def __init__(self):
        self._attr = '_%s_%d' % (self.__class__.__name__, self._next_id())
        
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return getattr(inst, self._attr)
        
    def __set__(self, inst, value):
        self.assert_valid(value)
        setattr(inst, self._attr, value)        
        
    def assert_valid(self, value):
        pass
    

class SymbolField(Field):
    _attr = '_symbol'
    
    def assert_valid(self, value):
        if not str.isupper(value):
            raise ValueError('symbol must be upper case, got %r' % value)
            

class PriceField(Field):
    _attr = '_price'
    
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError('price must be float, got %s' % type(value))
            
        if value <= 0:
            raise ValueError('price most be positive, got %s' % price)
        
    
class Stock:
    symbol = SymbolField()
    price = PriceField()
    low = PriceField()
    
    def __init__(self, symbol, price, low):
        self.symbol = symbol
        self.price = price
        self.low = low
        
    def __repr__(self):
        cls = self.__class__.__name__
        return '%s(%r, %r, %r)' % (cls, self.symbol, self.price, self.low)

brka = Stock('BRK.A', 216298.00, 216297.00)
brka


# In[160]:


vars(brka)


# In[172]:


class Math:
    @staticmethod
    def neg(val):
        return -val
    
Math.neg(10)


# In[173]:


m = Math()
m.neg(10)


# In[220]:


class StaticMethod:
    def __init__(self, func):
        self.func = func
        
    def __get__(self, inst, owner):
        return self.func
 
    
class Math:
    @StaticMethod
    def neg(val):
        return -val


# In[221]:


m = Math()
m.neg(10)


# In[222]:


Math.neg(10)


# In[201]:


class Point:
    def __init__(self, x, y):
        self.x, self.y = x, y
        
    @classmethod
    def from_str(cls, val):
        """From string in format 'x,y'"""
        x, y = map(float, val.split(','))
        return cls(x, y)
    
    def __repr__(self):
        return '%s(%r, %r)' % (self.__class__.__name__, self.x, self.y)
    
p = Point.from_str('2.3, 3.4')
p


# In[223]:


from functools import partial

class ClassMethod:
    def __init__(self, func):
        self.func = func
        
    def __get__(self, inst, owner):
        return partial(self.func, owner)
    
    
class Point:
    def __init__(self, x, y):
        self.x, self.y = x, y
        
    @ClassMethod
    def from_str(cls, val):
        """From string in format 'x,y'"""
        # print(cls)
        x, y = map(float, val.split(','))
        return cls(x, y)
    
    def __repr__(self):
        return '%s(%r, %r)' % (self.__class__.__name__, self.x, self.y)

p = Point.from_str('2.3, 3.4')
p


# In[224]:


p.from_str('1.1,2.2')


# In[225]:


class Person:
    def __init__(self, first, last):
        self.first = first
        self.last = last
        
    @property
    def name(self):
        return '%s %s' % (self.first, self.last)
    
p = Person('Tim', 'Peters')
p.name


# In[227]:


class Property:
    def __init__(self, func):
        self.func = func
        
    def __get__(self, inst, value):
        if not inst:
            return self
        
        return self.func(inst)

class Person:
    def __init__(self, first, last):
        self.first = first
        self.last = last
        
    @Property
    def name(self):
        return '%s %s' % (self.first, self.last)
    
p = Person('Tim', 'Peters')
p.name


# In[1]:


class Person:
    def __init__(self, name):
        self.name = name
        
    def greet(self):
        return "Hi, I'm %s. How are you?" % self.name
    
p = Person('Elvis')
p.greet()


# In[2]:


p.greet.__get__


# In[3]:


meth = p.greet.__get__(p, Person)
meth()


# In[ ]:




