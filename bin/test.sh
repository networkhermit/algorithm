#!/bin/bash

set -e

if (( $# > 2 )); then
    exit
fi

if [ -z "${1}" ]; then
    ROOT=.
elif [ -d "${1}" ]; then
    ROOT=${1}
else
    exit
fi

testingPhaseOfC () {
    echo '---------------------------[ C ]--------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.c" | sort); do
        fname=$(basename "${instance}" .c)

        if gcc   -O0 -g3 -Wall -Werror -std=c11 -x c -o "${fname}" "${instance}" -lm; then
            ./"${fname}"
            rm -f "${fname}"
        fi

        if clang -O0 -g3 -Wall -Werror -std=c11 -x c -o "${fname}" "${instance}" -lm; then
            ./"${fname}"
            rm -f "${fname}"
        fi

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfCPP () {
    echo '---------------------------[ C++ ]------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.cpp" | sort); do
        fname=$(basename "${instance}" .cpp)

        if g++     -O0 -g3 -Wall -Werror -std=c++14 -x c++ -o "${fname}" "${instance}"; then
            ./"${fname}"
            rm -f "${fname}"
        fi

        if clang++ -O0 -g3 -Wall -Werror -std=c++14 -x c++ -o "${fname}" "${instance}"; then
            ./"${fname}"
            rm -f "${fname}"
        fi

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfGo () {
    echo '---------------------------[ Go ]-------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.go" | sort); do
        fname=$(basename "${instance}" .go)

        if go run "${instance}"; then
            go clean -cache
        fi

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfJava () {
    echo '---------------------------[ Java ]-----------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.java" | sort); do
        fname=$(basename "${instance}" .java)

        if javac -d local_build -g -deprecation -Werror -Xlint:all,-path "${instance}"; then
            java -cp local_build -ea "${fname}"
            rm -rf local_build
        fi

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfLisp () {
    echo '---------------------------[ Lisp ]-----------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.lisp" | sort); do
        fname=$(basename "${instance}" .lisp)

        sbcl --noinform --non-interactive --no-sysinit --userinit ~/.fun/minimal.lisp \
            --load "${instance}" --eval "(when (fboundp 'main) (main))"

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfNodeJS () {
    echo '---------------------------[ NodeJS ]---------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.js" | sort); do
        fname=$(basename "${instance}" .js)

        node --throw-deprecation --trace-warnings "${instance}"

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfPHP () {
    echo '---------------------------[ PHP ]------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.php" | sort); do
        fname=$(basename "${instance}" .php)

        php -e "${instance}"

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfPython () {
    echo '---------------------------[ Python ]---------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.py" | sort); do
        fname=$(basename "${instance}" .py)

        python3 -d -W:all -B "${instance}"

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

testingPhaseOfRuby () {
    echo '---------------------------[ Ruby ]-----------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.rb" | sort); do
        fname=$(basename "${instance}" .rb)

        ruby -w "${instance}"

        echo -e "\\t\\t---------------------------    [${fname}]"
    done
}

echo '=============== TESTING AUTOMATION -- [STARTED] ================'

case "${2}" in
    '')
        testingPhaseOfC
        testingPhaseOfCPP
        testingPhaseOfGo
        testingPhaseOfJava
        testingPhaseOfLisp
        testingPhaseOfNodeJS
        testingPhaseOfPHP
        testingPhaseOfPython
        testingPhaseOfRuby
        ;;
    c)
        testingPhaseOfC
        ;;
    c++)
        testingPhaseOfCPP
        ;;
    go)
        testingPhaseOfGo
        ;;
    java)
        testingPhaseOfJava
        ;;
    lisp)
        testingPhaseOfLisp
        ;;
    node)
        testingPhaseOfNodeJS
        ;;
    php)
        testingPhaseOfPHP
        ;;
    python)
        testingPhaseOfPython
        ;;
    ruby)
        testingPhaseOfRuby
        ;;
    -)
        time testingPhaseOfC
        time testingPhaseOfCPP
        time testingPhaseOfGo
        time testingPhaseOfJava
        time testingPhaseOfLisp
        time testingPhaseOfNodeJS
        time testingPhaseOfPHP
        time testingPhaseOfPython
        time testingPhaseOfRuby
        ;;
    *)
        exit
        ;;
esac

echo '=============== TESTING AUTOMATION -- [FINISHED] ==============='
