# Title
SQLite - The Serialization Format You Need To Use

# Brief Summary 
> If your proposal is accepted this will be made public and printed in the program. Should be one paragraph, maximum 400 characters.

SQLite is an amazing product. It has types, schema (sort of), a query language and you can store up to 281 terabytes in a single file.
In this talk we see how you can use SQLite from Pandas/Python, learn some best practices and see some common pitfalls.

# Brief Bullet Point Outline
> Brief outline. Will be made public if your proposal is accepted. Edit using Markdown.

- Why SQLite?
- Using SQLite from Pandas
- Fun with types
- Faster inserts

# Description 
> Detailed outline. Will be made public if your proposal is accepted. Edit using Markdown.

At first, we'll have a look at amazing SQLite, it has billions of installs, can hold up to 281 terabytes in a single file and supports SQL (doh!)

Then we'll see how you can use SQLite from pandas for storing and retrieving data. We'll look into storing data in SQLite and how to query the database using SQL. We'll also talk about transactions and how they will save you some day.

Then we'll move to some oddities of SQLite. We'll see how to read `TIMESTAMP` using `sqlite3.PARSE_DECLTYPES`, how to enforce types with `CHECK` constraints.

Finally, insert data faster by dropping indices and using WAL.
