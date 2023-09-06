# A 10% Discount

import pandas as pd

df = pd.DataFrame([
    ['Bugs', True, 72.3],
    ['Daffy', False, 30.7],
    ['Tweety', True, 23.5],
    ['Elmer', False, 103.9],
], columns=['Customer', 'Member', 'Amount'])

# Give members 10% discount
df[df['Member']]['Amount'] *= 0.9
print(df)
