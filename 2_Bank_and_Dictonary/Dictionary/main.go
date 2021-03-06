package main

import (
	"fmt"

	"github.com/TheIncredibleMulk/learngo/2_Bank_and_Dictonary/Dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}

/*
err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search("hello")
	fmt.Println(hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
*/
