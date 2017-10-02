%.html: %.py
	pygmentize -O full,style=vs -f html $< > $@ 

html: asm.html config.html fib.html macro.html

asm.html: asm.py

config.html: config.py

test:
	./asm.py fib.asm | ./sim.py -R0 10

test2:
	./asm2.py fib.asm | ./sim.py -R0 10
