# Language

## Navigation

Panzer provides the `cd` keyword for path navigation:

```sh
~ > cd ..
/home > cd ..
/ >
/ > cd -
/home > cd ^
/ > cd ^
/home > cd ^
/~ >
```

## Matching filenames

Panzer provides the ability to perform matches on filenames:

```sh
~ > ls Down*
test.pdf
~ > ls ?ownload*
test.pdf
```

## Running commands

Panzer uses the `$PATH` variable to lookup the location of the executable we
want to run:

```sh
~ > ecko "test"
exec: "ecko": executable file not found in $PATH
~ > echo "test"
test
```

## Pipe

> **WARN**
>
> This is still a WIP feature, implementation, api and documentation is subject
> to frequent change

Redirect the output of one command into an other using the `|`:

```sh
~ > ls -la | grep "panzer"
```

## Environment variables

Use `set <VAR> <value>` to define a variable, use `env` to inspect the currently defined environment variables:

```sh
~ > set EDITOR nvim
~ > set PROMPT "\u \d $ "
~ > env
EDITOR=nvim
PROMPT="\u \d $ "
```

Defining variables with multiple arguments joins them with `:`:

```sh
~ > set TEST a b c
~ > echo $TEST
a:b:c
```

Variables in variable definitions are allowed:

```sh
~ > set PATH $PATH $HOME/.config/go/bin
```

Variable can be embedded everywhere:

```sh
~ > ls /home/$USER/.panzerc
# xnacly's panzer config
# https://github.com/xnacly

# --- env ---
set PROMPT "\u \d $ "
set EDITOR "nvim"
~ > echo $HOME
/home/teo
```

## Aliases

> **WARN**
>
> This is still a WIP feature, implementation, api and documentation is subject
> to frequent change

```sh
~ > alias gs "git status"
~ > alias gc "git commit -m"
~ > alias
alias gs "git status"
alias gc "git commit -m"
```

Aliases can be used as commands:

```sh
~ > gc "lang/preprocessor: reworked variable embedding"
```
