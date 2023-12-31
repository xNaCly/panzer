# Prompt

## Placeholders

Panzer supports prompt customization via placeholders in the `PROMPT`
environment variable.

A placeholder is a character prefixed with a backslash, such as `\u` for the
name of the currently logged in user.

Currently the following placeholders are supported:

| Placeholder | Description                          | Computation time                 |
| ----------- | ------------------------------------ | -------------------------------- |
| `\u`        | name of the currently logged in user | shell startup                    |
| `\h`        | name of the currently active host    | shell startup                    |
| `\w`        | current path / `$PWD` env variable   | directory change & shell startup |
| `\d`        | name of the current directory        | directory change & shell startup |
| `\D`        | current date                         | shell startup                    |
| `\t`        | current time (24h) `15:04:05`        | directory change & shell startup |
| `\s`        | shell name                           | shell startup (`panzer`)         |
| `\T`        | current time (12h) `03:04:05PM`      | directory change & shell startup |
| `\U`        | current time (unixepoch)             | directory change & shell startup |
| `\b`        | git branch name                      | every command & shell startup    |
| `\S`        | git status, `M` for modified         | every command & shell startup    |

The following colors are also supported:

| Placeholder | Color          | Escape code |
| ----------- | -------------- | ----------- |
| `\0`        | Reset          | `\033[0m`   |
| `\1`        | Red            | `\033[31m`  |
| `\2`        | Green          | `\033[32m`  |
| `\3`        | Yellow         | `\033[33m`  |
| `\4`        | Bright Red     | `\033[91m`  |
| `\5`        | Bright Green   | `\033[92m`  |
| `\6`        | Bright Yellow  | `\033[93m`  |
| `\7`        | Bright Blue    | `\033[94m`  |
| `\8`        | Bright Magenta | `\033[95m`  |
| `\9`        | Bright Cyan    | `\033[96m`  |

### Example:

```gpnzr
teo@comfyputer ~/programming/panzer master > set PROMPT "\u %"
teo %
```

## Prompt control environment variables

Enable any of the following environment variables by setting them to `1`.

| Environment variable | Description    | Effect                                          |
| -------------------- | -------------- | ----------------------------------------------- |
|                      | default prompt | `teo@comfyputer ~/programming/panzer master > ` |
