from nlp import tokenize

text = 'In the face of ambiguity, refuse the temptation to guess.'

for _ in range(10_000):
    tokenize(text)
