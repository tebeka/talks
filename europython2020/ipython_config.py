# Configuration file for ipython.

# lines of code to run at IPython startup.
c.InteractiveShellApp.exec_lines = [
    '%autoreload 2',
]

# A list of dotted module names of IPython extensions to load.
c.InteractiveShellApp.extensions = [
    'autoreload',
]

c.TerminalInteractiveShell.confirm_exit = False

c.TerminalInteractiveShell.editing_mode = 'vi'
c.TerminalInteractiveShell.prompt_includes_vi_mode = False

c.AliasManager.user_aliases = [
   ('cow', 'echo %l | cowsay'),
   ('mb', 'echo 💣 | cowsay'),
   ('wiifm', 'echo "WIIFM" | cowsay'),
   ('show', 'xdg-open %l'),
]
