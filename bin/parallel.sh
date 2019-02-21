#!/bin/bash

set -e

if [ ! -d "${PLT}" ]; then
    exit
fi

for i in ../codes/Algorithms/*/* ../codes/DataStructures/*; do
    (
    cd "${i}"
    mkdir -p Implementation Test
    rm -rf Implementation/* Test/*
    find -L "${PLT}" -type f -name "$(basename "${PWD}")"'.*' -exec ln -sv {} Implementation/ \;
    find -L "${PLT}" -type f -name "$(basename "${PWD}")"'Test.*' -exec ln -sv {} Test/ \;
    )
done
