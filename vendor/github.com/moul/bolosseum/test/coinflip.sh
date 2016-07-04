#!/bin/sh

echo "${0}: << ${1}" >&2

send() {
    echo "${0}: >> ${@}" >&2
    echo "${@}"
}

action="$(echo ${1} | jq -r .action)"

case $action in
    init)
        send '{"name": "bot1"}'
        ;;
    play-turn)
        if [ "$(( ( RANDOM % 2 ) ))" = "1" ]; then
            send '{"play": "head"}'
        else
            send '{"play": "ship"}'
        fi
        ;;
    *)
        send '{"error": "not implemented action: '${action}'"}'
        ;;
esac
