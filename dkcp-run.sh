#!/bin/sh
cd "$(dirname "$0")" || exit

print_horizontal_line() {
    printf '%*s\n' "${COLUMNS:-$(tput cols)}" '' | tr ' ' =
}

err_msg() { echo "$@" ;} >&2
usage() {
    echo "\033[0;91m"
    print_horizontal_line
    err_msg "Usage:"
    err_msg "  ./dkcp-run.sh [--dev-build <service>] [--log] [COMMAND]"
    err_msg "  ./dkcp-run.sh --help"
    err_msg ""
    err_msg "Options:"
    err_msg "  --dev-build [service]    Run 'dkcp up -d --build' command."
    err_msg "                           'all' for running all or write specific service."
    err_msg "  --log                    Show docker-compose logs. Place this at the last."
    err_msg "  --staging"
    err_msg ""
    err_msg "  --start [service]        Start container."
    err_msg "                           'all' for running all or write specific service."
    err_msg "  --stop [service]         Stop container."
    err_msg "                           'all' for running all or write specific service."
    err_msg "  --down [-v]              Remove container. Use '-v' to remove related volumes."
    print_horizontal_line
    echo "\033[0m"
    exit 1
}

echo ""
echo "\033[0;96mWebpractice - Run containers\033[0m"
echo ""

if [ "$#" -eq 0 ]; then
    usage && exit 1
fi

while true; do
    if [ "$#" -eq 0 ]; then exit 1; fi
    case "$1" in
        --dev-build)
            shift
            case $1 in (-*|"") err_msg "--dev-build: Wrong argument!"; usage; esac
            if [ "$1" = "all" ]; then
                echo >&2 "Develop - build all services."
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build
            else
                echo >&2 "Develop - build $1"
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build "$1"
            fi
            shift; continue
            ;;
        --start)
            shift
            case $1 in (-*|"") err_msg "--start: Wrong argument!"; usage; esac
            if [ "$1" = "all" ]; then
                echo >&2 "Start all containers."
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml start
            else
                echo >&2 "Start $1"
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml start "$1"
            fi
            shift; continue
            ;;
        --stop)
            shift
            case $1 in (-*|"") err_msg "--stop: Wrong argument!"; usage; esac
            if [ "$1" = "all" ]; then
                echo >&2 "Stop all containers."
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml stop
            else
                echo >&2 "Stop $1"
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml stop "$1"
            fi
            shift; continue
            ;;
        --down)
            shift
            if [ "$1" = "-v" ]; then
                echo >&2 "Down containers and volumes."
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml down -v
            else
                echo >&2 "Down containers."
                docker-compose -f docker-compose.yml -f docker-compose.dev.yml down
            fi
            shift; continue
            ;;
        --staging)
            # # 도커레지스트리(private? public?) 에서 풀 한뒤 deploy
            echo "staging"
            shift; continue
            ;;
        --log)
            docker-compose logs -f
            shift; continue
            ;;
        --help)
            usage && exit 1
            ;;
        *)
            usage && exit 1
            ;;
    esac
done

# References
# https://mug896.github.io/bash-shell/getopts.html