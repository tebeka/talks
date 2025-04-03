import tkinter as tk
from tkinter import ttk

# Calculator GUI:
# ____________
# ____________
# [7][8][9][/]
# [4][5][6][*]
# [1][2][3][-]
# [0][.][C][+]
# [    =     ]

buttons = (
    ('7', '8', '9', '/'),
    ('4', '5', '6', '*'),
    ('1', '2', '3', '-'),
    ('0', '.', 'C', '+'),
)

root = tk.Tk()
root.title('Calc')
tk.Grid.rowconfigure(root, 0, weight=1)
tk.Grid.columnconfigure(root, 0, weight=1)
display = ttk.Entry(root)
display.pack(fill=tk.X)

fr = ttk.Frame(root)
for r, row in enumerate(buttons, 1):
    for c, text in enumerate(row, 1):
        if text == 'C':
            command = lambda:  display.delete(0, tk.END)
        else:
            command = lambda: display.insert(tk.END, text)
        b = ttk.Button(fr, text=text, command=command)
        b.grid(row=r, column=c, sticky='NSWE')
fr.pack(expand=True)


def calc():
    expr = display.get()
    try:
        val = eval(expr)
    except (SyntaxError, NameError):
        return
    display.delete(0, tk.END)
    display.insert(0, str(val))


ttk.Button(root, text='=', command=calc).pack(fill=tk.X)
root.mainloop()
