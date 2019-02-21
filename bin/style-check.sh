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

styleCheckPhaseC () {
    echo '---------------------------[ C ]--------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.c" | sort); do
        fname=$(basename "${instance}" .c)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '(\s*)\s*{$' "${instance}"; then
            echo '>>>> MISSING PARAMETER LIST <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done

    for instance in $(find -L "${ROOT}" -type f -name "*.h" | sort); do
        fname=$(basename "${instance}" .h)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '(\s*)\s*{$' "${instance}"; then
            echo '>>>> MISSING PARAMETER LIST <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhaseCPP () {
    echo '---------------------------[ C++ ]------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.cpp" | sort); do
        fname=$(basename "${instance}" .cpp)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if ! grep -q '\<using\s*namespace\s*std\>' "${instance}"; then
            echo '>>>> NOT USING DEFAULT NAMESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done

    for instance in $(find -L "${ROOT}" -type f -name "*.hpp" | sort); do
        fname=$(basename "${instance}" .hpp)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<using\s*namespace\>' "${instance}"; then
            echo '>>>> NOT USING DEFAULT NAMESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhaseGo () {
    echo '---------------------------[ Go ]-------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.go" | sort); do
        fname=$(basename "${instance}" .go)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q ';$' "${instance}"; then
            echo '>>>> REDUNDANT SEMICOLON <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<if\s*(.*)\s*{$' "${instance}"; then
            echo '>>>> REDUNDANT PARENTHESIS <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<for\s*(.*)\s*{$' "${instance}"; then
            echo '>>>> REDUNDANT PARENTHESIS <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhaseJava () {
    echo '---------------------------[ Java ]-----------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.java" | sort); do
        fname=$(basename "${instance}" .java)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '[^[:upper:]>]>>[^>]' "${instance}"; then
            echo '>>>> NOT USING LOGICAL RIGHT SHIFT <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhaseLisp () {
    echo '---------------------------[ Lisp ]-----------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.lisp" | sort); do
        fname=$(basename "${instance}" .lisp)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhaseNodeJS () {
    echo '---------------------------[ NodeJS ]---------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.js" | sort); do
        fname=$(basename "${instance}" .js)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q ';$' "${instance}"; then
            echo '>>>> REDUNDANT SEMICOLON <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '[^>]>>[^>]' "${instance}"; then
            echo '>>>> NOT USING LOGICAL RIGHT SHIFT <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '[^!=][!=]=[^=]' "${instance}"; then
            echo '>>>> NOT USING STRICT (IN)EQUALITY <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhasePHP () {
    echo '---------------------------[ PHP ]------------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.php" | sort); do
        fname=$(basename "${instance}" .php)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhasePython () {
    echo '---------------------------[ Python ]---------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.py" | sort); do
        fname=$(basename "${instance}" .py)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q ';$' "${instance}"; then
            echo '>>>> REDUNDANT SEMICOLON <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<\(el\)\?if\s*(.*)\s*:$' "${instance}"; then
            echo '>>>> REDUNDANT PARENTHESIS <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<while\s*(.*)\s*:$' "${instance}"; then
            echo '>>>> REDUNDANT PARENTHESIS <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

styleCheckPhaseRuby () {
    echo '---------------------------[ Ruby ]-----------------------------'

    for instance in $(find -L "${ROOT}" -type f -name "*.rb" | sort); do
        fname=$(basename "${instance}" .rb)

        if grep -q '\s$' "${instance}"; then
            echo '>>>> TRAILING WHITESPACE <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q ';$' "${instance}"; then
            echo '>>>> REDUNDANT SEMICOLON <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<\(els\)\?if\s*(.*)$' "${instance}"; then
            echo '>>>> REDUNDANT PARENTHESIS <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi

        if grep -q '\<while\s*(.*)$' "${instance}"; then
            echo '>>>> REDUNDANT PARENTHESIS <<<<'
            echo -e "\\t\\t---------------------------    [${fname}]"
        fi
    done
}

echo '================= STYLE CHECKING -- [STARTED] =================='

case "${2}" in
    '')
        styleCheckPhaseC
        styleCheckPhaseCPP
        styleCheckPhaseGo
        styleCheckPhaseJava
        styleCheckPhaseLisp
        styleCheckPhaseNodeJS
        styleCheckPhasePHP
        styleCheckPhasePython
        styleCheckPhaseRuby
        ;;
    c)
        styleCheckPhaseC
        ;;
    c++)
        styleCheckPhaseCPP
        ;;
    go)
        styleCheckPhaseGo
        ;;
    java)
        styleCheckPhaseJava
        ;;
    lisp)
        styleCheckPhaseLisp
        ;;
    node)
        styleCheckPhaseNodeJS
        ;;
    php)
        styleCheckPhasePHP
        ;;
    python)
        styleCheckPhasePython
        ;;
    ruby)
        styleCheckPhaseRuby
        ;;
    -)
        time styleCheckPhaseC
        time styleCheckPhaseCPP
        time styleCheckPhaseGo
        time styleCheckPhaseJava
        time styleCheckPhaseLisp
        time styleCheckPhaseNodeJS
        time styleCheckPhasePHP
        time styleCheckPhasePython
        time styleCheckPhaseRuby
        ;;
    *)
        exit
        ;;
esac

echo '================= STYLE CHECKING -- [FINISHED] ================='
