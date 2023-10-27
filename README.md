# gopnzr

gopnzr - a minimal, fast and fun to use shell written in go

> _Warning_
> gopnzr is in active development

## Features

- [x] command execution, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#Command-Search-and-Execution)
- [x] filename expansion, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#Filename-Expansion)
- [ ] Redirection, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#Redirections)
- [ ] Piping, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#Pipelines)
- [x] director navigation, `cd`, see [bash docs](https://www.gnu.org/software/bash/manual/html_node/Bourne-Shell-Builtins.html#index-cd)
- [ ] extensive help system via build-in `help [topic]` command
- [ ] background / asynchronous tasks, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#Lists)
- [x] self managed directory list, similar to `pushd`, `popd` - but populated with the last visited directories, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#The-Directory-Stack)
- [x] customizable prompt, see [bash docs](https://www.gnu.org/software/bash/manual/bash.html#Controlling-the-Prompt)
- [x] minimal configuration file
- [x] no posix or bash compatibility - none!
- [x] history list
- [x] variables
- completion engine
  - [x] available files (in current dir)
  - [ ] available commands
  - [ ] available cli flags `ls <option/flag>`
  - [ ] available man pages, `man <page>`
