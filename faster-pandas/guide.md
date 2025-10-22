# Performance Guidlines

## CPU

- "for loop" is the enemy, use vectorized operation
- Prefer pandas or numpy functions on built-in ones
- Caching is awesome
- Do it cheaper

## Memory

- Load only what you need
- Arrow rocks
- Use right size types
- Can you do it in chunks 
