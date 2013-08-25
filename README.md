# Regex Matcher
Brian Kernighan wrote about how Rob Pike implement a beautiful and simple regular expression matcher in C.

After discussing how complex existing regular expression packages were for their book Rob disappeared into his office appeared again in no more than an hour or two with the 30 lines of C code that subsequently appeared in Chapter 9 of TPOP. That code implements a regular expression matcher that handles these constructs:

    c    matches any literal character c
    .    matches any single character
    ^    matches the beginning of the input string
    $    matches the end of the input string
    *    matches zero or more occurrences of the previous character

This is a re-implementation of the same in a language Rob Pike later himself designed.

Full story: [http://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html](http://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html)