![cdd](./logo.jpg)

# cdd

**English** | [日本語](./README_JA.md)

[![test](https://github.com/koooyooo/cdd/actions/workflows/test.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/test.yaml)
[![lint](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml/badge.svg)](https://github.com/koooyooo/cdd/actions/workflows/lint.yaml)

Jump to bookmarked directories with short aliases — a smarter `cd`.

Reach deep paths by name or index. The `cdd` binary resolves the path; your shell’s `builtin cd` does the move (shell integration required).

```bash
# Before
$ cd ../../Documents/projects/cdd/chart

# After
$ cdd chart
```

## Install

```bash
go install github.com/koooyooo/cdd@latest
```

Enable shell integration (zsh example):

```bash
# Interactive: prints the hook and optionally copies the eval line to the clipboard
cdd init zsh

# Or append to your rc file
echo 'eval "$(cdd init zsh)"' >> ~/.zshrc
```

For bash, use `cdd init bash` and `~/.bashrc`.

To complete alias names in zsh, run `eval` after `compinit`:

```bash
autoload -Uz compinit && compinit
eval "$(cdd init zsh)"
```

Bash completes alias names and subcommands on `cdd <Tab>` with no extra setup.

## Usage

Run `$ cdd {command}`. With no arguments, `cdd` prints help.

### Jump (with shell integration)

Pass an alias name or a `list` index to `cd` into that directory.

```bash
$ cdd docs
$ pwd
/Users/me/Documents

$ cdd 1
$ pwd
/Users/me/Documents
```

### Commands

| Command | Alias | Description |
| --- | --- | --- |
| `print` | `p` | Resolve an alias to an absolute path on stdout |
| `init` | — | Print shell integration for bash or zsh |
| `list` | — | List registered aliases |
| `add` | — | Add an alias |
| `remove` | `rm` | Remove an alias |
| `move-up` / `move-down` | `up` / `down` | Reorder aliases in the list |
| `edit` | — | Open the config file in your editor |

#### `print` / `p`

Resolve an alias name or index to an absolute path. Used by shell integration and scripts.

```bash
$ cdd print docs
/Users/me/Documents

$ cd "$(cdd print docs)"
```

#### `init`

Print bash/zsh shell integration to stdout. Besides the jump function, it registers completion for alias names and subcommands. The argument after `remove`, `move-up`, `move-down`, and `print` completes as an alias name; the argument after `init` completes as `bash` or `zsh`; the path argument to `add` completes as a filename. When stdout is a TTY, you are prompted (Y/n) to copy the `eval "$(cdd init …)"` line to the clipboard for pasting into your rc file.

```bash
$ eval "$(cdd init zsh)"
$ cdd init zsh
```

#### `list`

List registered aliases. On first run, `home` and `docs` are registered by default.

```bash
$ cdd list
    0 | home | ${HOME}
    1 | docs | ${HOME}/Documents
```

#### `add`

Register an alias with `$ cdd add <path>` (alias = last directory name), `$ cdd add <name> <path>`, or `$ cdd add <path> -a <name>`. `<path>` may be absolute or relative.

```bash
$ cdd add /Users/me/Downloads           # alias = Downloads
$ cdd add . -a work                     # current directory with explicit alias
$ cdd add dls /Users/me/Downloads
$ cdd add docs '${HOME}/Documents'      # quote ${HOME} to prevent shell expansion
```

#### `remove` / `rm`

Remove an alias with `$ cdd remove <name>`. You can also use a `list` index.

```bash
$ cdd remove dls
```

#### `move-up` / `move-down`

Move an alias up or down in the list. Defaults to one step; pass a second argument for the distance. `up` / `down` remain as aliases.

```bash
$ cdd move-up dls
$ cdd move-up dls 2
$ cdd move-down dls 2
```

#### `edit`

Open the config file in your default editor (not supported on Windows yet).

```bash
$ cdd edit
```

## Config

Aliases are stored in `${HOME}/.cdd.yaml`. You can edit this file directly.

```yaml
- name: home
  dir: ${HOME}
- name: docs
  dir: ${HOME}/Documents
```

`dir` should be an absolute path. As an exception, paths relative to `${HOME}` are also accepted. Other environment variables are not expanded.
