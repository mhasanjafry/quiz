Solution:

I am going to use Patricia tree for my approach, as it gives a space optimization storage of strings, and you can search on it in 0(m*N) complexity, where m is length of string, and N is size of alphabet. I am chosing the following implementation of Patricia Tree (github.com/tchap/go-patricia/patricia trie)

I will make two passes on the dataset. First, I will insert all strings into the tree, and the second time, I will test each string, if it is composed of any compound words. There can be two approaches to this: Either  (A) read the file and store it in any array for second pass, or (B) read the file again from the file system for second pass. I assume that in a distributed system, file I/O might be very expensive, and storing the file in memory might be a cheap option, so I have chosen option (A). 

In second pass, every string is checked if its length is more than current longest compound string. If it is, the string is traversed, iterating through the characters, and recursively choosing a substring of characters, and checking if the corresponding substring's suffix is a string present in the tree or a compound word itself.

How to run?

Install the patricia tree dependency: go get -u github.com/tchap/go-patricia/patricia

Then execute command "go build solution.go". Next execute command "./solution [optional filename]"