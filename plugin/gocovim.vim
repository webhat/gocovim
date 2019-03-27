if !exists("g:gocovim_command")
    let g:gocovim_command = "gocover"
endif

function! GoCoVim()
    silent !clear
    execute "!" . g:gocovim_command 
    sign unplace 1
    source .cadre/coverage.vim
endfunction

nnoremap <buffer> <leader>r :call GoCoVim()<cr>
