# Write serval programs with three levels of nesting and include FOR loops, WHILE loops, and IF statements. 
# Before you run your program try to predict what your program is going to print. 
# If the result is different from your prediction try to figure out why.

# Count down 10
count=10
while [[ $count -gt 0 ]]
do
    echo "count is $count"
    let count=$count-1
done

for number in {1..3}
do
    for letter in {a..b}
    do
        echo "number is $number, letter is $letter"
    done
done
