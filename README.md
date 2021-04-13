# Scrambled-strings
It's a solution to find out how many of the words from a dictionary file appear as substrings in a long string of
characters either in their original form or in their scrambled form.

### Description
To run this assignment we need to have two file which will be used to fetch dictionary words and input scrambled strings.
So for this program we two files
1. Dictionary file with multiple words on new line and each word on new line.

Ex: dictionary.txt
```
axpaj
apxaj
dnrbt
pjxdn
abd
```

2. We need scrambled strings input file which one is used to compare with the dictionary word's appearance.

Ex. inputs.txt
```
aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt
```
These strings can be in scrambled form.

## To run this program:
You need to clone this repo on your local in the `/home/{username}/go/src/`
directory
```
git clone https://github.com/sanvidhans/scrambled-strings
```

#### Change the directory
```
cd scrambled-strings
```

#### Run the unit test case:
```
go test ./...
```

#### Build the program:
```
go build -o scrambled-strings .
```

#### To run this program you need to pass dictionary file and input file path in the arguments:
```
./scrambled-strings --dictionary [DICTIONARY_FILE_PATH] --input [INPUT_FILE_PATH]
```

example:
```
./scrambled-strings --dictionary sampleFiles/dictionary.txt --input sampleFiles/inputs.txt 
```

Here in this repo we have sample dictionary and input file formats
so you can try this sample files or you can try your own file path.


# Sample Output
## Input Dictionary:
filepath: sampleFiles/dictionary.txt
```
axpaj
apxaj
dnrbt
pjxdn
abd
dnrbt
dw
kasma
aplpe
apple
mnago
```

## Input Strings
```
aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt
kjhfdgmnghjsdtrklituwedjannbdslkamsahdo
kituryendmahdjflfmahdkappleghunghatmangoluccy
```

## Output:
```
Case#1: 4
Case#2: 1
Case#3: 3
```
Here in the output of the program you get the number of count the dictionary words appeared on the line number of scrambled string.

`Case#1: 4` means the 4 words from dictionary file found at the line number 1.
so there are 3 scrambled strings on separate line in the inputs.txt file.
and out of that the program could find the:
- 4 words(scrambled form or original form) on line number 1.
- 1 word on line number 2.
- 3 words on the line number 3.




Note:
```
Dockerfile file implementation is not done in this program.
it is inprogress.
so for now dockerfile may not works
```
