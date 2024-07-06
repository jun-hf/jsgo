6/07
Today I managed to build the lexer for my programming language. I have decided to represent token as an integer to keep the size of each token small. It is the lexer jobs to correctly get each token and the corresponding literal of the token. On top of that I have added support for floating number token in my lexer. 

I have also added errors on to the lexer struct. The design decision here is to ensure that there are errors reporting in the lexer and the parser.

For the lexer, the errors represent the errors by characters, wherelse for in parsers it is reponsible fo checking errors in terms grammers

The next step is to see how does go, does positioning and then I can starting to port some the features into my language.
