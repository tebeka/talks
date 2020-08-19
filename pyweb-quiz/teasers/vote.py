# Let's Vote

import re

text = 'The vote was 21 in favour, 13 against and 10 abstentions'
match = re.search(r'(\d+).*(\d+).*(\d+)', text)
print(match.group(1), match.group(2), match.group(3))
