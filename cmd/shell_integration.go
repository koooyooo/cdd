package cmd

import (
	"fmt"
	"strings"
)

// completionCommandWords are offered on the first argument, together with alias names.
var completionCommandWords = []string{
	"list", "l",
	"add", "a",
	"remove", "rm", "delete", "del",
	"move-up", "up", "move-down", "down",
	"edit",
	"print", "p",
	"init",
	"help",
	"completion",
	"-h", "--help",
}

// aliasArgWords take an alias name as the next argument.
var aliasArgWords = []string{
	"remove", "rm", "delete", "del",
	"move-up", "up", "move-down", "down",
	"print", "p",
}

var pathArgWords = []string{"add", "a"}

var initShellWords = []string{"bash", "zsh"}

const shellFunctionTmpl = `
cdd() {
  if [ $# -eq 0 ]; then
    command cdd
    return
  fi
  case "$1" in
    %s)
      command cdd "$@"
      return
      ;;
  esac
  dir="$(command cdd print "$@")" || return
  builtin cd -- "$dir"
}
`

const bashCompletionTmpl = `
_cdd_add_matches() {
  local cur="$1"
  shift
  local word
  for word in "$@"; do
    if [ -z "${cur}" ] || [ "${word#"${cur}"}" != "${word}" ]; then
      COMPREPLY+=("${word}")
    fi
  done
}

_cdd_aliases() {
  local cur="$1"
  local name
  while IFS= read -r name; do
    [ -n "${name}" ] || continue
    if [ -z "${cur}" ] || [ "${name#"${cur}"}" != "${name}" ]; then
      COMPREPLY+=("${name}")
    fi
  done < <(command cdd complete-aliases 2>/dev/null)
}

_cdd_files() {
  local cur="$1"
  local path
  while IFS= read -r path; do
    COMPREPLY+=("${path}")
  done < <(compgen -f -- "${cur}")
}

_cdd_complete() {
  local cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=()
  if [ "${COMP_CWORD}" -eq 1 ]; then
    _cdd_aliases "${cur}"
    _cdd_add_matches "${cur}" %s
    return
  fi
  if [ "${COMP_CWORD}" -eq 2 ]; then
    case "${COMP_WORDS[1]}" in
      %s)
        _cdd_aliases "${cur}"
        ;;
      init)
        _cdd_add_matches "${cur}" %s
        ;;
    esac
    return
  fi
  if [ "${COMP_CWORD}" -eq 3 ]; then
    case "${COMP_WORDS[1]}" in
      %s)
        _cdd_files "${cur}"
        ;;
    esac
  fi
}
complete -F _cdd_complete cdd
`

const zshCompletionTmpl = `
_cdd() {
  local -a aliases subs
  aliases=("${(@f)$(command cdd complete-aliases 2>/dev/null)}")
  aliases=("${(@)aliases:#}")
  subs=(%s)
  if (( CURRENT == 2 )); then
    if (( ${#aliases} )); then
      compadd -- "${aliases[@]}"
    fi
    compadd -- "${subs[@]}"
    return 0
  fi
  case "${words[2]}" in
    %s)
      if (( ${#aliases} )); then
        compadd -- "${aliases[@]}"
      fi
      ;;
    init)
      compadd -- %s
      ;;
    %s)
      if (( CURRENT == 4 )); then
        _files
      fi
      ;;
  esac
  return 0
}
if whence compdef >/dev/null 2>&1; then
  compdef _cdd cdd
fi
`

func shellIntegration(shell string) (string, error) {
	fn := fmt.Sprintf(shellFunctionTmpl, passthroughPattern())
	var comp string
	switch shell {
	case "bash":
		comp = fmt.Sprintf(
			bashCompletionTmpl,
			shellQuoteWords(completionCommandWords),
			strings.Join(aliasArgWords, "|"),
			shellQuoteWords(initShellWords),
			strings.Join(pathArgWords, "|"),
		)
	case "zsh":
		comp = fmt.Sprintf(
			zshCompletionTmpl,
			shellQuoteWords(completionCommandWords),
			strings.Join(aliasArgWords, "|"),
			shellQuoteWords(initShellWords),
			strings.Join(pathArgWords, "|"),
		)
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
	return strings.TrimPrefix(fn, "\n") + comp, nil
}

func passthroughPattern() string {
	words := append([]string{}, completionCommandWords...)
	words = append(words, "complete-aliases")
	return strings.Join(words, "|")
}

func shellQuoteWords(words []string) string {
	quoted := make([]string, len(words))
	for i, w := range words {
		quoted[i] = "'" + strings.ReplaceAll(w, "'", `'\''`) + "'"
	}
	return strings.Join(quoted, " ")
}
