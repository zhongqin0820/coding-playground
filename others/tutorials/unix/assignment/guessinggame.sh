#!/usr/bin/env bash
function ask {
    echo "Hey, guess how many files are in the current directory..."
    read -p "Please input an integer: " guess
    if [[ ! $guess =~ ^[0-9]+$ ]]; then
        echo "Your input is illegal!"
        ask
    fi
}

function main {
    files=$(ls | wc -l)
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