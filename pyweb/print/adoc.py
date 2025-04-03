# https://asciidoctor.org/docs/asciidoc-writers-guide/#description-lists


def description_list(title: str, desc: dict[str, list[str]]):
    return '\n'.join(_descrpition_list(title, desc))


def _descrpition_list(title, desc):
    yield f'.{title}'
    for section, products in desc.items():
        yield f'{section}::'
        for item in products:
            yield f'* {item}'


shopping_list = {
    'Dairy': ['Milky', 'Cheese'],
    'Produce': ['Lemon', 'Mango'],
    'Bakery': ['Baguette'],
}


print(description_list('Shopping List', shopping_list))

# .Shopping List
# Dairy::
# * Milky
# * Cheese
# Produce::
# * Lemon
# * Mango
# Bakery::
# * Baguette
