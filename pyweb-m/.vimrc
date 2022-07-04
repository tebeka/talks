" Wheels are zip files
au BufReadCmd *.whl call zip#Browse(expand("<amatch>"))
