# Why Vim is Awesome

Given at [PyWeb IL](https://www.meetup.com/PyWeb-IL/events/257305259/)

[Slides](https://docs.google.com/presentation/d/1-7DR3vsiv8GEgUall8X50EhuEDQKCxmVEioVZb5CBPY/edit)

## Playbook
- [Lindy effect](https://en.wikipedia.org/wiki/Lindy_effect)
    - vi 1976; 42 years ago
    - vim 1991; 27 years ago
- it's everywhere
- jokes then move to vim
- car analogy
- math2
- start without file
    - show exit instructions
- open tests/test_sqrt.py
- no mouse (find image?)
- dual mode
- movement
- CTRL-] on sqrt
    - plug
    - jedi plugin
    - \r (val to value)
    - :Gdiff
- import
    - show ale (C-N)
- goto tests/sqrt_cases.json
    - add space
    - gf (goto file)
- %!jq .
- back to test and :Gdiff
- :make dist
    - Show Makefile
- :NT
    - open dist/
    - open wheel
- C-p / NERDTree
    - fzf in the command line
- add MANIFEST.in 
    - `recursive-include tests *.py`
- open .gitignore
- add math2.egg-info/
    - math<CTRL-X F> 
- [airline](https://github.com/vim-airline/vim-airline)
    - git branch, number of words
- setup.py
    - gx
- snippets
- open __init__.py
- def<TAB> .. (add)
    - C-B
- plugins
    - [PyCharm](https://plugins.jetbrains.com/plugin/164-ideavim)
    - [VSCode](https://marketplace.visualstudio.com/items?itemName=vscodevim.vim)
    - bash/zsh: `set -o vi`
    - ~/.inputrc: `set editing-mode vi`
    - IPython: `c.TerminalInteractiveShell.editing_mode = 'vi'`
    - [vimium](https://chrome.google.com/webstore/detail/vimium/dbepggeogbaibhgnhhndojpepiihcmeb?hl=en)
- scriptable with Python
    - :help python
- learn
    - vim adventures
    - vim golf
    - https://www.openvim.com/
- gang sign
