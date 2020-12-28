# vim

  vim-modes: `normal`, `command`, `insert`, `visual`, `select`
    `visual` vs `select`: 
      作用: 主要是不同模式下，键盘映射不一样。可以对visual mode 多一套键盘映射 
      转换: `visual` <C-g> -> `select`
  格式: `{cmd} [args] {lhs} {rhs}`
    **args**: <silent><buffer><expr>...
    **cmd**:
      | 命 令                    | 模 式                        |
      | :map   :noremap  :unmap  | 普通、可视、选择、操作符等待 |
      | :nmap  :nnoremap :nunmap | 普通                         |
      | :vmap  :vnoremap :vunmap | 可视与选择                   |
      | :smap  :snoremap :sunmap | 选择                         |
      | :xmap  :xnoremap :xunmap | 可视                         |
      | :omap  :onoremap :ounmap | 操作符等待                   |
      | :map!  :noremap! :unmap! | 插入与命令行                 |
      | :imap  :inoremap :iunmap | 插入                         |
      | :lmap  :lnoremap :lunmap | 插入、命令行、Lang-Arg       |
      | :cmap  :cnoremap :cunmap | 命令行                       |
      | :tmap  :tnoremap :tunmap | 终端作业                     |
