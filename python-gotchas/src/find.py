haystack = 'hello there'
needle = 'bugs'
if haystack.find(needle):
    print('WAT?')
# => WAT
needle = 'hello'
if not haystack.find(needle):
    print('is broken')
# => is broken
