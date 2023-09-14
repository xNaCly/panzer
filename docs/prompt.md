# Prompt

## Placeholders

gpnzr supports prompt customization via placeholders in the `GPNZR_PROMPT`
environment variable.

A placeholder is a character prefixed with a backslash, such as `\u` for the
name of the currently logged in user.

Currently the following placeholders are supported:

| Placeholder | Description                                    |
| ----------- | ---------------------------------------------- |
| `\u`        | name of the currently logged in user           |
| `\h`        | name of the currently active host              |
| `\w`        | name of the current path / `$PWD` env variable |
| `\d`        | name of the current directory                  |

### Example:

```gpnzr
teo@comfyputer /home/teo/programming/gopnzr :: export GPNZR_PROMPT="\u"
teo %
```

## Prompt control environment variables

Enable any of the following environment variables by setting them to `1`.

| Environment variable     | Description                                  | Effect                                             |
| ------------------------ | -------------------------------------------- | -------------------------------------------------- |
|                          | default prompt                               | `teo@comfyputer   /home/teo/programming/gopnzr ::` |
| `GPNZR_PROMPT_SHORT_PWD` | enables shortening of paths to one character | `teo@comfyputer /h/t/p/g ::`                       |
