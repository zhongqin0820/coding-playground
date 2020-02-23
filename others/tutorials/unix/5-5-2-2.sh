# Write a bash script where you define two arrays inside of the script
# and the sum of the lengths of the arrays are printed to the console when the script is run.

# Step 1: Define two arrays
arr1=(a b c d e)
arr2=(1 2 3 4 5 6)
# Step 2: Output the sum
echo "len(arr1) + len(arr2) = $(expr ${#arr1[*]} + ${#arr2[*]})"
