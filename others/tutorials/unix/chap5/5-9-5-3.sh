# Write a program called range that takes one number as an argument and prints all of the numbers between that number and 0.
function range {
    local fn
    if [[ $1 -lt 0 ]]
    then
        fn=$(eval echo {$1..0})
    else
        fn=$(eval echo {0..$1})
    fi
    echo "## $fn"
}

echo "----------- TEST CASE -----------"
echo "range 6"
range 6
echo "## 0 1 2 3 4 5 6"
echo "----------- TEST CASE -----------"
echo "range -3"
range -3
echo "## -3 -2 -1 0"

# Write a program called extremes which prints the maximum and minimum values of a sequence of numbers.
function extremes {
    local max=$1
    local min=$1
    for num in $@ 
    do
        if [[ $max -lt $num ]]
        then
            let max=$num
        fi
        if [[ $min -gt $num ]]
        then
            let min=$num
        fi
    done
    echo "## $min $max"
}

echo "----------- TEST CASE -----------"
echo "extremes 8 2 9 4 0 3"
extremes 8 2 9 4 0 3
echo "## 0 9"
