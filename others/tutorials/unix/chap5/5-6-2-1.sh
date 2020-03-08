# Create 100 text files using brace expansion.
touch $(eval echo file{001..100}.txt)
