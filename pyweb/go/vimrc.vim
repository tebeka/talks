call plug#begin('~/.vim/plugged')

Plug 'SirVer/ultisnips'
Plug 'bling/vim-airline'
Plug 'cespare/vim-toml'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'davidhalter/jedi-vim'
Plug 'dcharbon/vim-flatbuffers'
Plug 'dyng/ctrlsf.vim'
Plug 'fatih/vim-go'
Plug 'flazz/vim-colorschemes'
Plug 'honza/vim-snippets'
Plug 'nvie/vim-flake8'
Plug 'peter-edge/vim-capnp'
Plug 'rust-lang/rust.vim'
Plug 'scrooloose/nerdcommenter'
Plug 'scrooloose/nerdtree'
Plug 'solarnz/thrift.vim'
Plug 'tmhedberg/matchit'
Plug 'tpope/vim-fugitive'
Plug 'tpope/vim-surround'
Plug 'vim-scripts/SWIG-syntax'
Plug 'vim-scripts/a.vim'
Plug 'w0rp/ale'

call plug#end()

" highlight python
let python_highlight_all=1

" syntastic
let g:syntastic_always_populate_loc_list = 1
let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 0


set nu
set ignorecase
set smartcase
set showmatch
set incsearch
set nohls

" Tab stuff
set tabstop=8
set softtabstop=4
set shiftwidth=4

" Use system clipboard
set clipboard=unnamedplus

" Make no noise
" set visualbell t_vb=
set noerrorbells

" airline
set laststatus=2

" GUI Stuff
if has("gui_running")
    " No toolbar or menu
    set guioptions=aAce
    colorscheme koehler
    set guifont=PT\ Mono\ 18
    " set guifont=PT\ Mono\ 20
endif

" Spelling
au BufNewFile,BufRead ChangeLog,*.txt,*.rst,*.markdown,*.md,*.yaml,*.yml setl spell

" Go
" let g:go_fmt_command = "goimports"
au BufNewFile,BufRead *.go setl noet
au BufNewFile,BufRead *.go setl ts=4
au BufNewFile,BufRead *.go setl tw=0

let g:go_auto_sameids = 1
let g:go_auto_type_info = 1
au FileType go nmap <leader>gt :GoDeclsDir<cr>

" .i are SWIG
au BufNewFile,BufRead *.i set ft=swig

" HTML
au BufNewFile,BufRead *.html setl ts=2
au BufNewFile,BufRead *.html setl sw=2
au BufNewFile,BufRead *.html setl tw=0

" HTML + jinja
au FileType htmljinja setl ts=2
au FileType htmljinja setl sw=2
au FileType htmljinja setl tw=0


func! ASCII_Clean()
    silent %! iconv -c -t ASCII
endfunc
comm! AC call ASCII_Clean()

" Space for page up/down
noremap <SPACE> <C-F>

" Explorer to show full info
let g:explDetailedList=1

" Use hidden (allow buffer switch when modified)
au BufNewFile,BufRead * set hidden

" Abbriviations
abb _3me_ Miki Tebeka <miki@353solutions.com>
abb _bp_ import pdb; pdb.set_trace()  # noqa
abb _cr_ # Copyright 353Solutions, All rights reserved
abb _date_ <C-R>=strftime("%Y-%m-%d")<CR>
abb _gme_ Miki Tebeka <miki.tebeka@gmail.com>
abb _ibp_ from IPython.core.debugger import Pdb; Pdb().set_trace()
abb _mme_ Miki Tebeka <miki@mikitebeka.com>
abb _plt_ import matplotlib.pyplot as plt
abb _py2_ #!/usr/bin/env python2
abb _py3_ #!/usr/bin/env python3
abb _py_ #!/usr/bin/env python
abb _rst_ .. vim: ft=rst spell
abb _sh_ #!/bin/bash

" Python stuff
au FileType python setl makeprg=flake8\ %
au FileType python syntax sync fromstart
au FileType python setl textwidth=79
au FileType python setl ts=4
autocmd BufWritePost,BufRead *.py call Flake8()
let g:flake8_show_in_gutter=1
let g:flake8_show_quickfix=0

" let g:jedi#auto_vim_configuration = 0
" let g:jedi#popup_select_first = 0
" set completeopt=menuone

" XML indent is two spaces
au FileType xml setl shiftwidth=2

" SQL indent is two spaces
au FileType sql setl shiftwidth=2

" Make file executable
func! MakeExecutable()
    w
    !chmod +x %
    e
endfunc
comm! MX silent call MakeExecutable()

" egg, wheel are just zip files (but can be directories as well)
au BufReadCmd *.egg,*.whl if !isdirectory(expand("<amatch>")) | 
            \   call zip#Browse(expand("<amatch>")) | endif


let NERDTreeIgnore = ['\.py[co]$', '__pycache__']
comm! NT NERDTree
comm! NTC NERDTreeClose
let g:NERDTreeWinSize = 18 

comm! Gr ! go run %
comm! Pr ! python %
comm! Pr2 ! python2 %

abb _ys_ **Youngstar:**
abb _gb_ **Graybeard:**

" ultisnips
let g:UltiSnipsExpandTrigger="<tab>"
let g:UltiSnipsJumpForwardTrigger="<c-b>"
let g:UltiSnipsJumpBackwardTrigger="<c-z>"

" clang
let g:clang_library_path='/usr/lib/libclang.so'
let g:clang_jumpto_declaration_key="<C-0>"

" java
autocmd FileType java setlocal omnifunc=javacomplete#Complete

" c++
autocmd FileType cpp setlocal et
autocmd FileType cpp setlocal sw=2
autocmd FileType cpp setlocal ts=2


" asm
au BufNewFile,BufRead *.ASM set ft=tasm
au BufNewFile,BufRead *.ASM set ts=2
au BufNewFile,BufRead *.ASM set sw=2

" Ignore
let g:ctrlp_custom_ignore = {
  \ 'dir':  '\.venv$\|\.vendor$',
  \ }

" Search
noremap <leader>s :CtrlSF 

let g:ale_sign_error = '⤫'
let g:ale_sign_warning = '⚠'
let g:airline#extensions#ale#enabled = 1
let g:ale_linters = {
\   'python': ['flake8'],
\}
