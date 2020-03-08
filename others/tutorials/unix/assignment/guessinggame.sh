#!/usr/bin/env bash
function ask {
    echo "Hey, guess how many files are in the current directory..."
    read guess
}

function main {
    files=$(ls -1 | wc -l)
    ask
    while [[ $guess -ne $files ]]; do
        if [[ $guess -lt $files ]]; then
            echo "Too low."
        else
            echo "Too high."
        fi
        ask
    done
    echo "Congratulations!"
}

main