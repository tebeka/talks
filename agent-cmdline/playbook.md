- CLI (and MCPs) is what your agent use to interact with the world
- Why Go?
    - No need for runtime
    - Easy to cross compile
    - Static executable
- basics of unix philosopy vs xkcd tar
- What's the difference
    - A lot of rules aren't new, but agents make it matter.
    - Agent
        - Tokens 
        - Structured output
        - Non interactive
- https://www.nyc.gov/site/tlc/about/tlc-trip-record-data.page
- make psql
- agents run a command without a terminal and read only captured output
- Usage & color issue
    go run ./0-initial query sql/sample.sql
    go run ./0-initial query sql/sample.sql | cat -v

    mention spinners ...
    comma in numbers
- TTY
    go run ./1-tty query sql/sample.sql
    go run ./1-tty query sql/sample.sql | cat -v

    errors/progress go to stderr
        - agent will try to parse "connecting to database", human will gloss over
        - agent cat do "cmd 2>/tmp/err.txt"
-  JSON
    go run ./2-json query sql/sample.sql
    go run ./2-json -format=json query sql/sample.sql

-  Interactive
    go run ./3-yes query sql/count.sql
    go run ./3-yes query -yes sql/count.sql

-  Schema to help agent
    go run ./4-schema schema
    go run ./4-schema schema trips

-  What to do next
    go run ./5-errors query sql/typo.sql
    go run ./5-errors query sql/bad_column.sql
    go run ./5-errors -format=json query sql/typo.sql

-  Bounded
    - kills context
    go run ./6-bounded query sql/big.sql
    go run ./6-bounded query -limit 3 sql/big.sql
    go run ./6-bounded -format=json query -limit 3 sql/big.sql
    go run ./6-bounded query -limit 0 sql/big.sql | tail -1
    - does not limit query on db

- security: read only user (joke slide)
    - the db enforces permissions, the cli can't

- https://duckdb.org/2026/10/09/agent-mode

##  TBD

- Keep query results on stdout. Send progress, warnings, and diagnostics to stderr.

---

slides
https://docs.google.com/presentation/d/1vzzbhxhZQMXw7ZcIWkf0935Bf15uViCDsfn52hhZMoM/

---

## DuckDB Agent Mode lessons

The DuckDB post describes agent mode for its CLI. Useful ideas for this talk are:

- Agent detection can use known environment variables together with stdout not being a terminal.
- Automatic behavior should stay out of the way when the user selects an explicit output format.
- Compact Markdown tables reduce formatting noise. Column types can be shown with names.
- Truncation must be explicit. DuckDB reports omitted rows and can stop large results early.
- Errors can use a compact JSON form on stderr.
- Schema inspection, query estimates, and progress can reduce guesswork.
- Give one short startup notice. Keep full instructions in help instead of repeating them on every call.

The post reports 59% fewer output tokens in its experiment, with all answers correct. It also reports no reduction in total cost or runtime. Use this as evidence that compact output can reduce output tokens, not as a claim that it always makes agent work cheaper or faster.
