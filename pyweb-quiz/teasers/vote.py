# Let's Vote

import re

text = 'The vote was 65 in favour, 43 against and 21 abstentions'
match = re.search(r'(\d+).*(\d+).*(\d+)', text)
print(match.group(1), match.group(2), match.group(3))
