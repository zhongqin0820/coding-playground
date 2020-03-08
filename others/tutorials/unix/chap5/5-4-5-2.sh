# Write a Bash script that takes one argument and prints “even” if the first argument is an even number 
# or “odd” if the first argument is an odd number.
if [[ $(expr $1 % 2) -eq 0 ]]
then
    echo "even"
else
    echo "odd"
fi
