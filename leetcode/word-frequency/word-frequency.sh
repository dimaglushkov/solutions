# Read from the file words.txt and output the word frequency list to stdout.
# one-line solution
xargs -a words.txt -n1 | sort | uniq -c | sort -r | awk '{print $2 " " $1}'


