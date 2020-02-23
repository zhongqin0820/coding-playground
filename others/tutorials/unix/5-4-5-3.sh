# Write a Bash script that takes two arguments. 
# If both arguments are numbers, print their sum, otherwise just print both arguments.
if [[ $1 =~ [0-9] ]] && [[ $2 =~ [0-9] ]]
then
    # res=$(echo "$1 + $2" | bc -l)
    res=$(expr $1 + $2)
    echo "$1 + $2 = $res"
else
    echo "$1 $2"
fi
