# Write a bash script where you define an array inside of the script, and the first argument for the script indicates the index of the array element that is printed to the console when the script is run.

# Step 1: Define an array
arr=(a b c d e f z)
# Step 2: Output the corresponding arr[index]
if [[ $1 -lt ${#arr[*]} ]]
then
    echo "The corresponding arr[$1] = ${arr[$1]}"
else
    echo "$1 is out of index, maximum index = $(expr ${#arr[*]} - 1)."
fi
