class Item:
    def __init__(self, sku, price, amount):
        self.sku = sku
        self.price = price
        self.amount = amount


class Cart:
    def __init__(self, items):
        self._items = items[:]

    def total(self):
        return sum(item.price * item.amount for item in self._items)


class BlackFridayCart(Cart):
    def __init__(self, items, discounted):
        super().__init__(items)
        self._items = discounted  # discounted items


items = [
    Item('MacMini', 2700, 1),
    Item('Dell Screen', 120, 1),
    Item('HDMI Cable', 7.5, 1),
]

discounted = {
    'Dell Screen': 100,
}

cart = BlackFridayCart(items, discounted)
