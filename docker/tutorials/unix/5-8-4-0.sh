# Write a function called plier which multiplies together a sequence of numbers.
function plier {
    local res=1
    for num in $@
    do 
        let res=$res\*$num
    done
    echo "## $res"
}
echo "----------- TEST CASE -----------"
echo "plier 7 2 3"
plier 7 2 3
echo "## 42"

# Write a function called isiteven that prints 1 if a number is even or 0 a number is not even.
function isiteven {
    if [[ $1%2 -eq 0 ]]
    then
        echo "## 1"
    else
        echo "## 0"
    fi
}
echo "----------- TEST CASE -----------"
echo "isiteven 42"
isiteven 42
echo "## 1"

# Write a function called nevens which prints the number of even numbers when provided with a sequence of numbers. Use isiteven when writing this function.
function nevens {
    local res=0
    for num in $@
    do
        if [[ $num%2 -eq 0 ]]
        then
            let res=$res+1
        fi
    done
    echo "## $res"
}
echo "----------- TEST CASE -----------"
echo "nevens 42 6 7 9 33"
nevens 42 6 7 9 33
echo "## 2"

# Write a function called howodd which prints the percentage of odd numbers in a sequence of numbers. Use nevens when writing this function.
function howodd {
    local res=0
    for num in $@
    do
        if [[ $num%2 -eq 1 ]]
        then
            let res=$res+1
        fi
    done
    echo "## $(echo "$res / $#" | bc -l )"
}
echo "----------- TEST CASE -----------"
echo "howodd 42 6 7 9 33"
howodd 42 6 7 9 33
echo "## .60"

# Write a function called fib which prints the number of fibonacci numbers specified.
function fib {
    local res
    f1=0
    f2=1
    Num=$@
    echo "The Fibonacci sequence for the number $Num is: "
    for (( i=0;i<Num;i++ ))
    do
        res+=($f1)
        fn=$((f1+f2))
        f1=$f2
        f2=$fn
    done
    echo "##${res[*]}"
}
echo "----------- TEST CASE -----------"
echo "fib 4"
fib 4
echo "## 0 1 1 2"
echo "----------- TEST CASE -----------"
echo "fib 10"
fib 10
echo "## 0 1 1 2 3 5 8 13 21 34"
