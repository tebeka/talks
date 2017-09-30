# slide
Hi, I'm Miki and I hate DSLs

# slide
Let's get there by the long route...

I was working at a company that made DVDs, doing tools (simulators, debuggers
...) for a custom chip. BTW: simulator was written in C++, started to add
Python extensions and made it faster :)

Visited a friend on a another team, they had no one
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

He said - whatever works, everything is better than this.

# slide
You get the idea, they wrote Python. And it took me 2 days to do it.

# slide
Toy assembly

# slide
Example encoding

# slide
Let's see code

# slide
Anyone can guess what this does?

# slide
fibonacci

# slide
Not far from Python bytecode
(clipped output)

# slide
Same assembly instruction can have several opcodes

# slide
Shift inside full instruction

# slide
program = list of instructions
labels = need to resolve in 2nd pass
instructions available for user

instruction decorator

# slide
register has a code
repr helps a lot when debugging

# slide
label knows it's location (one static file)

# slide
base class
line info
add to program
ABC for bits

# slide
Continue ASM
code
slot_bits - depend if immediate or register

# slide
str for output
got to love f-strings

# slide
JMP concrete instruction 
has a target
calling `__init__` will add to program

# slide
JMP cont
handle labels
again repr

# slide
JMPE just different opcode

# slide
MOV opcode depend on src

# slide
compile is super easy

# slide
command line

# slide
example output, LMK if you find a bug :)

# slide
few months later - macros?

# slide
already there

# slide
how it this connected?

# slide
language designer
tooling
document

more mental overload to users

# slide
Don't do DSL
dumb, stupid language

# slide
Python is dynamic enough for most of your needs

# slide
another example - config

# slide
json (no comments), yaml (OK), ini (no types), TOML (OK), DSL (no)
Was in a company that used about 5 different formats for configuration

# slide
why not python?

# slide
I use this in several project, super easy and works

# slide
demo invocation

# slide
security

# slide
SQL has built in eval

# slide
you need to be careful with Python - yaml.safe_load

# slide
trust no one

# slide
before you write your own DSL

# slide
Just use Python

# slide
questions

code contains a simulator as well
