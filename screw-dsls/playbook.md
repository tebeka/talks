# slide
Hi, I'm Miki and I hate DSLs

# slide
Let's get there by the long way...

I was working at a company that made DVDs, doing tools (simulators, debuggers
...) for a custom chip. Visited a friend on a another team, they had no one
working on tools for them so they wrote in machine code.

BTW: Anyone recognizes this? This is the Python DLL on my machine :)

He asked me if I can write an assembler for them. I said I'll talk to my boss,
which of course said I don't have time. I told him - 2 days, and he said if
it's 2 days - go for it.

# slide
So I want back to my friend and said, I can do it in two days, but the assembly
will be a bit different from what you're used to.

    MOV(AX, BX) -> MOV(AX, BX)
    LOOP1: -> LABEL('LOOP1')

He said - whatever works.
