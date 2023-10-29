# completion command
Generate the autocompletion script for redovc for the specified shell. This can be piped to a shell script for inclusion in a login rc file.

See each sub-command's help for details on how to use the generated script

## Usage
`renotevc completion [shellname]`

## Aliases
  None

## Examples
`renotevc completion > .renotevc_completion.sh`

## Available Commands
```
bash        Generate the autocompletion script for bash
fish        Generate the autocompletion script for fish
powershell  Generate the autocompletion script for powershell
zsh         Generate the autocompletion script for zsh
```

## Flags
`-h, --help   help for completion`

## Global Flags:
`-f, --firebase   Run commands using the firebase service`