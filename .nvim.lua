vim.keymap.set("n", "<leader>fw", ":tabe | terminal make<CR>:file RUN<CR>G")

vim.keymap.set("n", "<leader>fq", ":tabe | buf RUN<CR>i<C-c>")
vim.keymap.set("n", "<leader>ft", ":!make templ<CR>")
vim.keymap.set("n", "<leader>t", ":!make test<CR>")
