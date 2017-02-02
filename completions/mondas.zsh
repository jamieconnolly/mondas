#compdef mondas

local cmds line ret=1 state

_arguments -C \
  '--help[Display help information]' \
  '--version[Display version information]' \
  '1: :->cmds' && ret=0

case $state in
  cmds)
    cmds=(${(f)"$(_call_program commands ${service} completions 2>/dev/null)"})
    _describe -t commands "${service} commands" cmds && ret=0
    ;;
esac

return ret
