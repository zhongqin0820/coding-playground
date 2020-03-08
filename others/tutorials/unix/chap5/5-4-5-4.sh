# Write a Bash script that prints “Thank Moses it’s Friday” if today is Friday.
# (Hint: take a look at the date program).
if [[ $(date +%u) -eq 5 ]]
then
    echo "Thank Moses it’s Friday"
else
    echo "Today is $(date +%A)."
fi
