## Talk Ideas

## Divide and Prosper

Don't tangle data layers. DB, Business, API

### RS24

Title:
All Your Schemas Are Belong To Us

Description:
You placed a simple API over your user store. It's simple - the client provides the user ID, you get it from the database, convert to JSON and done! One more JIRA ticket to close. Congratulations! You just created both a huge technical debt and a security hole.

By tangling the database schema, with the business object and the API layer, you make it very hard to make schema changes. Also, if someone adds the user authentication token to the database, it'll automatically go out to any client.

In this talk, I'll make a claim that you should keep the database, business and API data models separate. Even at the cost of code duplication.


Outline:

[2 min] About me
[5 min] The problem
[7 min] Schemas and where they hurt
[7 min] A little copying is better than a little dependency
[5 min] Security considerations
[3 min] Conclusion

Audience Takeaway:
Use different schema & objects for each layer of your application.



## Code that writes code

For example stringer in Go.
Fixed parser (RDH example)
Mention generics
Generate `_ import` for config

## Get Comfortable

How being in your comfort zone is better. More productive, less pressure. Drive analogy - focus on road, not on car.

## Prototyping

- https://www.youtube.com/watch?v=vLxX3yZmw5Q
- Plan to throw one away ...

## Mutation, Rebinding & Scope
- Mutation vs rebinding
- LCGB
- namespaces = dicts all the way down globals, locals
- Hack: adding variable to bulitins
- global, nonlocal

## Fallacies of distributed computing
https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing

See how to overcome them in Go/Python

## Mise en place

How to have everything in place *before* you start codign
- RFC
- Ontology
- Use cases

## Open source a project
Based on the daily go

## Rest
How to rest, in dev, code, life ...
- Hammock driven development
- Python zen: Sparse is better than dense

## Faster Pandas

Vectorization, ufuncs, append, eval & query, join(?), sql & hdf5, numba, cython

## Choose a baseline

- Algorithms scoring (fraud: say no, trading: buy & hold ...)
    - https://en.wikipedia.org/wiki/Cohen%27s_kappa
- Performance - dumb, brute force ...
- https://medium.com/microsoftazure/9-advanced-tips-for-production-machine-learning-6bbdebf49a6f
- recursion base
- napkin math

## So you think you can print?

Go (https://blog.gopheracademy.com/advent-2018/fmt/)
Python: `__repr__`, `__format__`, `{name:<20}`

## Metaclasses and why you shouldn't use them
- What are they
- Implement ABC

## Testing Conference Banner
- QR code
- Image processing

