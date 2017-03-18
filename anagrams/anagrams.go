package main

import (
	"bufio"
	"fmt"
	"github.com/alexey-malov/go_projects/anagrams/anagrams"
	"log"
	"os"
)

func readLn(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}

func readTwoStrings(reader *bufio.Reader) (string, string, error) {
	readString := func(prompt string) (string, error) {
		fmt.Print(prompt)
		return readLn(reader)
	}

	s1, err := readString("Enter the first string: ")
	if err != nil {
		return "", "", fmt.Errorf("Failed to read first string: %s", err.Error())
	}

	s2, err := readString("Enter the second string: ")
	if err != nil {
		return "", "", fmt.Errorf("Failed to read second string: %s", err.Error())
	}

	return s1, s2, nil
}

func main() {
	fmt.Println("The program determines if one string is an anagram of another string.")

	reader := bufio.NewReader(os.Stdin)

	s1, s2, err := readTwoStrings(reader)
	if err != nil {
		log.Panicf(err.Error())
	}

	if anagrams.IsAnagram(s1, s2) {
		fmt.Printf("'%s' is an anagram of '%s'\n", s1, s2)
	} else {
		fmt.Printf("'%s' is not an anagram of '%s'", s1, s2)
	}
}
